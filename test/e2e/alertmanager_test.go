// Copyright 2019 The Cluster Monitoring Operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/openshift/cluster-monitoring-operator/pkg/manifests"
	"github.com/pkg/errors"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

func TestAlertmanagerVolumeClaim(t *testing.T) {
	err := f.OperatorClient.WaitForStatefulsetRollout(&v1beta2.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alertmanager-main",
			Namespace: f.Ns,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	cm := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cluster-monitoring-config",
			Namespace: f.Ns,
		},
		Data: map[string]string{
			"config.yaml": `alertmanagerMain:
  volumeClaimTemplate:
    spec:
      storageClassName: gp2
      resources:
        requests:
          storage: 2Gi
`,
		},
	}

	if err := f.OperatorClient.CreateOrUpdateConfigMap(cm); err != nil {
		t.Fatal(err)
	}

	var lastErr error
	// Wait for persistent volume claim
	err = wait.Poll(time.Second, 5*time.Minute, func() (bool, error) {
		_, err := f.KubeClient.CoreV1().PersistentVolumeClaims(f.Ns).Get("alertmanager-main-db-alertmanager-main-0", metav1.GetOptions{})
		lastErr = errors.Wrap(err, "getting alertmanager persistent volume claim failed")
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		if err == wait.ErrWaitTimeout && lastErr != nil {
			err = lastErr
		}
		t.Fatal(err)
	}

	err = f.OperatorClient.WaitForStatefulsetRollout(&v1beta2.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alertmanager-main",
			Namespace: f.Ns,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlertmanagerTrustedCA(t *testing.T) {
	var (
		factory = manifests.NewFactory("openshift-monitoring", nil)
		newCM   *v1.ConfigMap
		lastErr error
	)

	// Wait for the new ConfigMap to be created
	err := wait.Poll(time.Second, 5*time.Minute, func() (bool, error) {
		cm, err := f.KubeClient.CoreV1().ConfigMaps(f.Ns).Get("alertmanager-trusted-ca-bundle", metav1.GetOptions{})
		lastErr = errors.Wrap(err, "getting new trusted CA ConfigMap failed")
		if err != nil {
			return false, nil
		}

		newCM = factory.HashTrustedCA(cm, "alertmanager")
		if newCM == nil {
			lastErr = errors.New("no trusted CA bundle data available")
			return false, nil
		}

		return true, nil
	})
	if err != nil {
		if err == wait.ErrWaitTimeout && lastErr != nil {
			err = lastErr
		}
		t.Fatal(err)
	}

	// Wait for the new hashed trusted CA bundle ConfigMap to be created
	err = wait.Poll(time.Second, 5*time.Minute, func() (bool, error) {
		_, err := f.KubeClient.CoreV1().ConfigMaps(f.Ns).Get(newCM.Name, metav1.GetOptions{})
		lastErr = errors.Wrap(err, "getting new CA ConfigMap failed")
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		if err == wait.ErrWaitTimeout && lastErr != nil {
			err = lastErr
		}
		t.Fatal(err)
	}

	// Get Alertmanager StatefulSet and make sure it has a volume mounted.
	err = wait.Poll(time.Second, 5*time.Minute, func() (bool, error) {
		ss, err := f.KubeClient.AppsV1().StatefulSets(f.Ns).Get("alertmanager-main", metav1.GetOptions{})
		lastErr = errors.Wrap(err, "getting Alertmanager StatefulSet failed")
		if err != nil {
			return false, nil
		}

		if len(ss.Spec.Template.Spec.Containers[0].VolumeMounts) == 0 {
			return false, errors.New("Could not find any VolumeMounts, expected at least 1")
		}

		for _, mount := range ss.Spec.Template.Spec.Containers[0].VolumeMounts {
			if mount.Name == "configmap-" + newCM.Name {
				return true, nil
			}
		}

		lastErr = fmt.Errorf("no volume %s mounted", newCM.Name)
		return false, nil
	})
	if err != nil {
		if err == wait.ErrWaitTimeout && lastErr != nil {
			err = lastErr
		}
		t.Fatal(err)
	}
}

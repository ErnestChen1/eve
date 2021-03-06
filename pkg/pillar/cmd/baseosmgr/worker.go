// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package baseosmgr

import (
	"fmt"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/worker"
	"github.com/lf-edge/eve/pkg/pillar/zboot"
)

const (
	workInstall = "install"
)

// installWorkDescription install work we feed into the worker go routine
type installWorkDescription struct {
	contentID string
	ref       string
	target    string
}

// AddWorkInstall create a Work job to install the provided image to the target path
func AddWorkInstall(ctx *baseOsMgrContext, key, ref, target string) {
	d := installWorkDescription{
		contentID: key,
		ref:       ref,
		target:    target,
	}
	// we do not care about the errors much
	_ = ctx.worker.Submit(worker.Work{Key: key, Kind: workInstall, Description: d})
	log.Infof("AddWorkInstall(%s) done", key)
}

// installWorker implementation of work.WorkFunction that installs an image to a particular location
func installWorker(ctxPtr interface{}, w worker.Work) worker.WorkResult {
	d := w.Description.(installWorkDescription)

	result := worker.WorkResult{
		Key:         w.Key,
		Description: d,
	}

	if d.target == "" {
		result.Error = fmt.Errorf("installWorker: unassigned destination partition for %s", d.ref)
		result.ErrorTime = time.Now()
		return result
	}

	log.Infof("installWorker to install %s to %s", d.ref, d.target)
	err := zboot.WriteToPartition(log, d.ref, d.target)
	log.Infof("installWorker DONE install %s to %s: err %s",
		d.ref, d.target, err)

	if err != nil {
		result.Error = err
		result.ErrorTime = time.Now()
	}
	return result
}

// processInstallWorkResult handle the work result that was an installation
func processInstallWorkResult(ctxPtr interface{}, res worker.WorkResult) error {
	ctx := ctxPtr.(*baseOsMgrContext)
	d := res.Description.(installWorkDescription)
	baseOsHandleStatusUpdateUUID(ctx, d.contentID)
	return nil
}

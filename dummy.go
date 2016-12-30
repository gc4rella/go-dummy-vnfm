package main

import (
	"bytes"
	"fmt"
	"github.com/mcilloni/go-openbaton/catalogue"
	log "github.com/sirupsen/logrus"
	"time"
)

type handl struct {
	*log.Logger
}

// ActionForResume uses the given VNFR and VNFCInstance to return a valid
// action for resume. NoActionSpecified is returned in case no such Action exists.
func (h handl) ActionForResume(vnfr *catalogue.VirtualNetworkFunctionRecord,
	vnfcInstance *catalogue.VNFCInstance) catalogue.Action {
	return catalogue.NoActionSpecified
}

// CheckInstantiationFeasibility allows the VNFM to verify if the VNF instantiation is possible.
func (h handl) CheckInstantiationFeasibility() error {
	return nil
}

func (h handl) Configure(vnfr *catalogue.VirtualNetworkFunctionRecord) (*catalogue.VirtualNetworkFunctionRecord, error) {
	time.Sleep(3 * time.Second)

	return vnfr, nil
}

func (h handl) HandleError(vnfr *catalogue.VirtualNetworkFunctionRecord) error {
	h.Errorf("Error for VNFR %s\n", vnfr.Name)
	return nil
}

func (h handl) Heal(vnfr *catalogue.VirtualNetworkFunctionRecord,
	component *catalogue.VNFCInstance, cause string) (*catalogue.VirtualNetworkFunctionRecord, error) {
	return vnfr, nil
}

// Instantiate allows to create a VNF instance.
func (h handl) Instantiate(vnfr *catalogue.VirtualNetworkFunctionRecord, scripts interface{},
	vimInstances map[catalogue.ID][]*catalogue.VIMInstance) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Infof("Instantiating VNFR %v\n", vnfr)

	vnfr.Configurations.Append(&catalogue.ConfigurationParameter{
		ConfKey: "new_key",
		Value:   "new_value",
	})

	time.Sleep(3 * time.Second)

	return vnfr, nil
}

// Modify allows making structural changes (e.g.configuration, topology, behavior, redundancy model) to a VNF instance.
func (h handl) Modify(vnfr *catalogue.VirtualNetworkFunctionRecord,
	dependency *catalogue.VNFRecordDependency) (*catalogue.VirtualNetworkFunctionRecord, error) {

	h.Debugf("VirtualNetworkFunctionRecord VERSION is: %d\n", vnfr.HbVersion)
	h.Debugf("VirtualNetworkFunctionRecord NAME is: %s\n", vnfr.Name)
	h.Debugf("Got dependency: %v\n", dependency)

	buf := bytes.NewBufferString("\n")

	for key, value := range dependency.Parameters {
		buf.WriteString(fmt.Sprintf("\t%s: %v\n", key, value.Parameters))
	}

	h.Debug("Parameters are: {%s}", buf.String())

	time.Sleep(3 * time.Second)
	return vnfr, nil
}

// Query allows retrieving a VNF instance state and attributes. (not implemented)
func (h handl) Query() error {
	return nil
}

func (h handl) Resume(vnfr *catalogue.VirtualNetworkFunctionRecord,
	vnfcInstance *catalogue.VNFCInstance,
	dependency *catalogue.VNFRecordDependency) (*catalogue.VirtualNetworkFunctionRecord, error) {

	h.Infof("resume on VNFR '%s' with ID: %v\n", vnfr.Name, vnfr.ID)

	return vnfr, nil
}

// Scale allows scaling (out / in, up / down) a VNF instance.
func (h handl) Scale(scaleInOrOut catalogue.Action,
	vnfr *catalogue.VirtualNetworkFunctionRecord,
	component catalogue.Component,
	scripts interface{},
	dependency *catalogue.VNFRecordDependency) (*catalogue.VirtualNetworkFunctionRecord, error) {

	h.Infof("%v on VNFR %s with ID %v\n", scaleInOrOut, vnfr.Name, vnfr.ID)

	time.Sleep(3 * time.Second)

	return vnfr, nil
}

// Start starts a VNFR.
func (h handl) Start(vnfr *catalogue.VirtualNetworkFunctionRecord) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Infof("start VNFR: %s\n", vnfr.Name)
	time.Sleep(3 * time.Second)
	return vnfr, nil
}

func (h handl) StartVNFCInstance(vnfr *catalogue.VirtualNetworkFunctionRecord,
	vnfcInstance *catalogue.VNFCInstance) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Infof("start VNFCInstance '%s' with ID: %v\n", vnfcInstance.Hostname, vnfcInstance.ID)

	return vnfr, nil
}

// Stop stops a previously created VNF instance.
func (h handl) Stop(vnfr *catalogue.VirtualNetworkFunctionRecord) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Infof("stop VNFR: %s\n", vnfr.Name)
	//time.Sleep(3 * time.Second)
	return vnfr, nil
}

func (h handl) StopVNFCInstance(vnfr *catalogue.VirtualNetworkFunctionRecord,
	vnfcInstance *catalogue.VNFCInstance) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Infof("stop VNFCInstance '%s' with ID: %v\n", vnfcInstance.Hostname, vnfcInstance.ID)

	return vnfr, nil
}

// Terminate allows terminating gracefully or forcefully a previously created VNF instance.
func (h handl) Terminate(vnfr *catalogue.VirtualNetworkFunctionRecord) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Debugln("RELEASE_RESOURCES")
	h.Infof("Releasing resources for VNFR: %s\n", vnfr.Name)
	h.Debugf("Version is: %d\n", vnfr.HbVersion)

	for _, event := range vnfr.LifecycleEvents {
		if event.Event == catalogue.EventRelease {
			for _, vdu := range vnfr.VDUs {
				h.Debugf("Removing vdu: %v\n", vdu)

				time.Sleep(3 * time.Second)
			}
		}
	}

	return vnfr, nil
}

// UpdateSoftware allows applying a minor / limited software update(e.g.patch) to a VNF instance.
func (h handl) UpdateSoftware(script *catalogue.Script,
	vnfr *catalogue.VirtualNetworkFunctionRecord) (*catalogue.VirtualNetworkFunctionRecord, error) {
	h.Infof("Update software with script %v on VNFR %s with ID %v\n", script, vnfr.Name, vnfr.ID)

	time.Sleep(3 * time.Second)

	return vnfr, nil
}

// UpgradeSoftware allows deploying a new software release to a VNF instance.
func (h handl) UpgradeSoftware() error {
	return nil
}

// UserData returns a string containing UserData.
func (h handl) UserData() string {
	return "#!/usr/bin/env sh"
}

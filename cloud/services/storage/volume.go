package storage

import (
	"errors"
	"fmt"
	"strings"

	infrastructurev1beta1 "github.com/outscale-dev/cluster-api-provider-outscale.git/api/v1beta1"
	tag "github.com/outscale-dev/cluster-api-provider-outscale.git/cloud/tag"
	osc "github.com/outscale/osc-sdk-go/v2"
)

const (
	maxIops = 13000
	maxSize = 14901
)

// ValidateIops check that iops is valid
func ValidateIops(iops int32) (int32, error) {
	if iops < maxIops {
		return iops, nil
	} else {
		return iops, errors.New("Invalid iops")
	}
}

// ValidateSize check that size is valid
func ValidateSize(size int32) (int32, error) {
	if size < maxSize {
		return size, nil
	} else {
		return size, errors.New("Invalid Size")
	}
}

// ValidateVolumeType check that volumeType is a valid volumeType
func ValidateVolumeType(volumeType string) (string, error) {
	switch {
	case volumeType == "standard" || volumeType == "gp2" || volumeType == "io1":
		return volumeType, nil
	default:
		return volumeType, errors.New("Invalid volumeType")
	}
}

// ValidateSubregionName check that subregionName is a valid az format
func ValidateSubregionName(subregionName string) (string, error) {
	switch {
	case strings.Contains(subregionName, "1a") || strings.Contains(subregionName, "1b") || strings.Contains(subregionName, "2a") || strings.Contains(subregionName, "2b"):
		return subregionName, nil
	default:
		return subregionName, errors.New("Invalid subregionName")
	}
}

type OscVolumeInterface interface {
	CreateVolume(spec *infrastructurev1beta1.OscVolume, volumeName string) (*osc.Volume, error)
	DeleteVolume(volumeId string) error
	GetVolume(volumeId string) (*osc.Volume, error)
	ValidateVolumeIds(volumeIds []string) ([]string, error)
}

// CreateVolume create machine volume
func (s *Service) CreateVolume(spec *infrastructurev1beta1.OscVolume, volumeName string) (*osc.Volume, error) {
	size := spec.Size
	subregionName := spec.SubregionName
	volumeType := spec.VolumeType
	var volumeRequest osc.CreateVolumeRequest
	if volumeType == "io1" {
		iops := spec.Iops
		volumeRequest = osc.CreateVolumeRequest{
			Iops:          &iops,
			Size:          &size,
			SubregionName: subregionName,
			VolumeType:    &volumeType,
		}
	} else {
		volumeRequest = osc.CreateVolumeRequest{
			Size:          &size,
			SubregionName: subregionName,
			VolumeType:    &volumeType,
		}
	}
	oscApiClient := s.scope.GetApi()
	oscAuthClient := s.scope.GetAuth()
	volumeResponse, httpRes, err := oscApiClient.VolumeApi.CreateVolume(oscAuthClient).CreateVolumeRequest(volumeRequest).Execute()
	if err != nil {
		fmt.Printf("Error with http result %s", httpRes.Status)
		return nil, err
	}
	resourceIds := []string{*volumeResponse.Volume.VolumeId}
	err = tag.AddTag("Name", volumeName, resourceIds, oscApiClient, oscAuthClient)
	if err != nil {
		fmt.Printf("Error with http result %s", httpRes.Status)
		return nil, err
	}
	volume, ok := volumeResponse.GetVolumeOk()
	if !ok {
		return nil, errors.New("Can not create volume")
	}
	return volume, nil
}

// GetVolume retrieve volume from volumeId
func (s *Service) GetVolume(volumeId string) (*osc.Volume, error) {
	readVolumesRequest := osc.ReadVolumesRequest{
		Filters: &osc.FiltersVolume{
			VolumeIds: &[]string{volumeId},
		},
	}
	oscApiClient := s.scope.GetApi()
	oscAuthClient := s.scope.GetAuth()
	readVolumesResponse, httpRes, err := oscApiClient.VolumeApi.ReadVolumes(oscAuthClient).ReadVolumesRequest(readVolumesRequest).Execute()
	if err != nil {
		fmt.Printf("Error with http result %s", httpRes.Status)
		return nil, err
	}
	volumes, ok := readVolumesResponse.GetVolumesOk()
	if !ok {
		return nil, errors.New("Can not get volume")
	}
	if len(*volumes) == 0 {
		return nil, nil
	} else {
		volume := *volumes
		return &volume[0], nil
	}
}

// DeleteVolume delete machine volume
func (s *Service) DeleteVolume(volumeId string) error {
	deleteVolumeRequest := osc.DeleteVolumeRequest{VolumeId: volumeId}
	oscApiClient := s.scope.GetApi()
	oscAuthClient := s.scope.GetAuth()
	_, httpRes, err := oscApiClient.VolumeApi.DeleteVolume(oscAuthClient).DeleteVolumeRequest(deleteVolumeRequest).Execute()
	if err != nil {
		fmt.Printf("Error with http result %s", httpRes.Status)
		return err
	}
	return nil
}

// ValidatePublicIpIds validate the list of id by checking each volume resource and return volume resource that currently exist
func (s *Service) ValidateVolumeIds(volumeIds []string) ([]string, error) {
	readVolumeRequest := osc.ReadVolumesRequest{
		Filters: &osc.FiltersVolume{
			VolumeIds: &volumeIds,
		},
	}
	oscApiClient := s.scope.GetApi()
	oscAuthClient := s.scope.GetAuth()
	readVolume, httpRes, err := oscApiClient.VolumeApi.ReadVolumes(oscAuthClient).ReadVolumesRequest(readVolumeRequest).Execute()
	if err != nil {
		fmt.Printf("Error with http result %s", httpRes.Status)
		return nil, err
	}
	var validVolumeIds []string
	volumes, ok := readVolume.GetVolumesOk()
	if !ok {
		return nil, errors.New("Can not get volume")
	}
	if len(*volumes) != 0 {
		for _, volume := range *volumes {
			volumeId := volume.GetVolumeId()
			validVolumeIds = append(validVolumeIds, volumeId)
		}
	}
	return validVolumeIds, nil
}
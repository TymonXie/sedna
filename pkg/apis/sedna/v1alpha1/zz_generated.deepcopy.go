// +build !ignore_autogenerated

/*
Copyright The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregationWorker) DeepCopyInto(out *AggregationWorker) {
	*out = *in
	out.Model = in.Model
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregationWorker.
func (in *AggregationWorker) DeepCopy() *AggregationWorker {
	if in == nil {
		return nil
	}
	out := new(AggregationWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregationWorkerSpec) DeepCopyInto(out *AggregationWorkerSpec) {
	*out = *in
	in.CommonWorkerSpec.DeepCopyInto(&out.CommonWorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregationWorkerSpec.
func (in *AggregationWorkerSpec) DeepCopy() *AggregationWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(AggregationWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigModel) DeepCopyInto(out *BigModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigModel.
func (in *BigModel) DeepCopy() *BigModel {
	if in == nil {
		return nil
	}
	out := new(BigModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudWorker) DeepCopyInto(out *CloudWorker) {
	*out = *in
	out.Model = in.Model
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudWorker.
func (in *CloudWorker) DeepCopy() *CloudWorker {
	if in == nil {
		return nil
	}
	out := new(CloudWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonWorkerSpec) DeepCopyInto(out *CommonWorkerSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParaSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonWorkerSpec.
func (in *CommonWorkerSpec) DeepCopy() *CommonWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(CommonWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dataset) DeepCopyInto(out *Dataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dataset.
func (in *Dataset) DeepCopy() *Dataset {
	if in == nil {
		return nil
	}
	out := new(Dataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dataset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetList) DeepCopyInto(out *DatasetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dataset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetList.
func (in *DatasetList) DeepCopy() *DatasetList {
	if in == nil {
		return nil
	}
	out := new(DatasetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatasetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetSpec) DeepCopyInto(out *DatasetSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetSpec.
func (in *DatasetSpec) DeepCopy() *DatasetSpec {
	if in == nil {
		return nil
	}
	out := new(DatasetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetStatus) DeepCopyInto(out *DatasetStatus) {
	*out = *in
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetStatus.
func (in *DatasetStatus) DeepCopy() *DatasetStatus {
	if in == nil {
		return nil
	}
	out := new(DatasetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployModel) DeepCopyInto(out *DeployModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployModel.
func (in *DeployModel) DeepCopy() *DeployModel {
	if in == nil {
		return nil
	}
	out := new(DeployModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploySpec) DeepCopyInto(out *DeploySpec) {
	*out = *in
	out.Model = in.Model
	in.Trigger.DeepCopyInto(&out.Trigger)
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	in.HardExampleMining.DeepCopyInto(&out.HardExampleMining)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploySpec.
func (in *DeploySpec) DeepCopy() *DeploySpec {
	if in == nil {
		return nil
	}
	out := new(DeploySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeWorker) DeepCopyInto(out *EdgeWorker) {
	*out = *in
	out.Model = in.Model
	in.HardExampleMining.DeepCopyInto(&out.HardExampleMining)
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeWorker.
func (in *EdgeWorker) DeepCopy() *EdgeWorker {
	if in == nil {
		return nil
	}
	out := new(EdgeWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvalSpec) DeepCopyInto(out *EvalSpec) {
	*out = *in
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvalSpec.
func (in *EvalSpec) DeepCopy() *EvalSpec {
	if in == nil {
		return nil
	}
	out := new(EvalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FLJobCondition) DeepCopyInto(out *FLJobCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FLJobCondition.
func (in *FLJobCondition) DeepCopy() *FLJobCondition {
	if in == nil {
		return nil
	}
	out := new(FLJobCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FLJobSpec) DeepCopyInto(out *FLJobSpec) {
	*out = *in
	in.AggregationWorker.DeepCopyInto(&out.AggregationWorker)
	if in.TrainingWorkers != nil {
		in, out := &in.TrainingWorkers, &out.TrainingWorkers
		*out = make([]TrainingWorker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FLJobSpec.
func (in *FLJobSpec) DeepCopy() *FLJobSpec {
	if in == nil {
		return nil
	}
	out := new(FLJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FLJobStatus) DeepCopyInto(out *FLJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]FLJobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FLJobStatus.
func (in *FLJobStatus) DeepCopy() *FLJobStatus {
	if in == nil {
		return nil
	}
	out := new(FLJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedLearningJob) DeepCopyInto(out *FederatedLearningJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedLearningJob.
func (in *FederatedLearningJob) DeepCopy() *FederatedLearningJob {
	if in == nil {
		return nil
	}
	out := new(FederatedLearningJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedLearningJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedLearningJobList) DeepCopyInto(out *FederatedLearningJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedLearningJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedLearningJobList.
func (in *FederatedLearningJobList) DeepCopy() *FederatedLearningJobList {
	if in == nil {
		return nil
	}
	out := new(FederatedLearningJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedLearningJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardExampleMining) DeepCopyInto(out *HardExampleMining) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParaSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardExampleMining.
func (in *HardExampleMining) DeepCopy() *HardExampleMining {
	if in == nil {
		return nil
	}
	out := new(HardExampleMining)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILDataset) DeepCopyInto(out *ILDataset) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILDataset.
func (in *ILDataset) DeepCopy() *ILDataset {
	if in == nil {
		return nil
	}
	out := new(ILDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILJobCondition) DeepCopyInto(out *ILJobCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILJobCondition.
func (in *ILJobCondition) DeepCopy() *ILJobCondition {
	if in == nil {
		return nil
	}
	out := new(ILJobCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILJobSpec) DeepCopyInto(out *ILJobSpec) {
	*out = *in
	out.Dataset = in.Dataset
	out.InitialModel = in.InitialModel
	in.TrainSpec.DeepCopyInto(&out.TrainSpec)
	in.EvalSpec.DeepCopyInto(&out.EvalSpec)
	in.DeploySpec.DeepCopyInto(&out.DeploySpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILJobSpec.
func (in *ILJobSpec) DeepCopy() *ILJobSpec {
	if in == nil {
		return nil
	}
	out := new(ILJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ILJobStatus) DeepCopyInto(out *ILJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ILJobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ILJobStatus.
func (in *ILJobStatus) DeepCopy() *ILJobStatus {
	if in == nil {
		return nil
	}
	out := new(ILJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncrementalLearningJob) DeepCopyInto(out *IncrementalLearningJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncrementalLearningJob.
func (in *IncrementalLearningJob) DeepCopy() *IncrementalLearningJob {
	if in == nil {
		return nil
	}
	out := new(IncrementalLearningJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IncrementalLearningJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncrementalLearningJobList) DeepCopyInto(out *IncrementalLearningJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IncrementalLearningJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncrementalLearningJobList.
func (in *IncrementalLearningJobList) DeepCopy() *IncrementalLearningJobList {
	if in == nil {
		return nil
	}
	out := new(IncrementalLearningJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IncrementalLearningJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialModel) DeepCopyInto(out *InitialModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialModel.
func (in *InitialModel) DeepCopy() *InitialModel {
	if in == nil {
		return nil
	}
	out := new(InitialModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JointInferenceService) DeepCopyInto(out *JointInferenceService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JointInferenceService.
func (in *JointInferenceService) DeepCopy() *JointInferenceService {
	if in == nil {
		return nil
	}
	out := new(JointInferenceService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JointInferenceService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JointInferenceServiceCondition) DeepCopyInto(out *JointInferenceServiceCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JointInferenceServiceCondition.
func (in *JointInferenceServiceCondition) DeepCopy() *JointInferenceServiceCondition {
	if in == nil {
		return nil
	}
	out := new(JointInferenceServiceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JointInferenceServiceList) DeepCopyInto(out *JointInferenceServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JointInferenceService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JointInferenceServiceList.
func (in *JointInferenceServiceList) DeepCopy() *JointInferenceServiceList {
	if in == nil {
		return nil
	}
	out := new(JointInferenceServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JointInferenceServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JointInferenceServiceSpec) DeepCopyInto(out *JointInferenceServiceSpec) {
	*out = *in
	in.EdgeWorker.DeepCopyInto(&out.EdgeWorker)
	in.CloudWorker.DeepCopyInto(&out.CloudWorker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JointInferenceServiceSpec.
func (in *JointInferenceServiceSpec) DeepCopy() *JointInferenceServiceSpec {
	if in == nil {
		return nil
	}
	out := new(JointInferenceServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JointInferenceServiceStatus) DeepCopyInto(out *JointInferenceServiceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JointInferenceServiceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]Metric, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JointInferenceServiceStatus.
func (in *JointInferenceServiceStatus) DeepCopy() *JointInferenceServiceStatus {
	if in == nil {
		return nil
	}
	out := new(JointInferenceServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Model) DeepCopyInto(out *Model) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Model.
func (in *Model) DeepCopy() *Model {
	if in == nil {
		return nil
	}
	out := new(Model)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Model) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelList) DeepCopyInto(out *ModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Model, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelList.
func (in *ModelList) DeepCopy() *ModelList {
	if in == nil {
		return nil
	}
	out := new(ModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpec) DeepCopyInto(out *ModelSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpec.
func (in *ModelSpec) DeepCopy() *ModelSpec {
	if in == nil {
		return nil
	}
	out := new(ModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelStatus) DeepCopyInto(out *ModelStatus) {
	*out = *in
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = (*in).DeepCopy()
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]Metric, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelStatus.
func (in *ModelStatus) DeepCopy() *ModelStatus {
	if in == nil {
		return nil
	}
	out := new(ModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParaSpec) DeepCopyInto(out *ParaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParaSpec.
func (in *ParaSpec) DeepCopy() *ParaSpec {
	if in == nil {
		return nil
	}
	out := new(ParaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmallModel) DeepCopyInto(out *SmallModel) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmallModel.
func (in *SmallModel) DeepCopy() *SmallModel {
	if in == nil {
		return nil
	}
	out := new(SmallModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timer) DeepCopyInto(out *Timer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timer.
func (in *Timer) DeepCopy() *Timer {
	if in == nil {
		return nil
	}
	out := new(Timer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainSpec) DeepCopyInto(out *TrainSpec) {
	*out = *in
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	in.Trigger.DeepCopyInto(&out.Trigger)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainSpec.
func (in *TrainSpec) DeepCopy() *TrainSpec {
	if in == nil {
		return nil
	}
	out := new(TrainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingWorker) DeepCopyInto(out *TrainingWorker) {
	*out = *in
	in.WorkerSpec.DeepCopyInto(&out.WorkerSpec)
	out.Dataset = in.Dataset
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingWorker.
func (in *TrainingWorker) DeepCopy() *TrainingWorker {
	if in == nil {
		return nil
	}
	out := new(TrainingWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingWorkerSpec) DeepCopyInto(out *TrainingWorkerSpec) {
	*out = *in
	in.CommonWorkerSpec.DeepCopyInto(&out.CommonWorkerSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingWorkerSpec.
func (in *TrainingWorkerSpec) DeepCopy() *TrainingWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(TrainingWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	if in.Timer != nil {
		in, out := &in.Timer, &out.Timer
		*out = new(Timer)
		**out = **in
	}
	out.Condition = in.Condition
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}
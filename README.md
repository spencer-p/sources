# Knative Eventing Sources

This was my intern project at Google in 2019 with the Knative Eventing team. At
the time, Knative was fleshing out their library of Pod-specable resources and
my work was to create a set of easy-to-use Custom Resource Definitions (CRDs) to
lower the barrier of entry to writing Sources for Knative.

A source is an object from which Cloud Events originates from; in this
repository it is referred to as a source regardless of whether or not the event
originates on the cluster or is imported from somewhere else.

I wrote three separate CRDs to extend common useful Kubernetes objects, demos
for each of them, and a sidecar that could be used to simplify writing code for
those objects.

[Technical details are can be found in /docs/spec/README.md](docs/spec/README.md).

# Debugging

Even after having read everything this guide has to offer it is possible that a
`HelmRelease` fails and you want to debug it to get to the cause. This may be
a bit harder at first than you were used to while working with just `helm`
because you are no longer in direct control but the Helm Operator is doing the
work for you.

This last section of the guide will give you some pointers on how to debug a
failing `HelmRelease` resource.

## Getting the reason of failure

If a release fails the reason of failure will be logged in the Helm Operator's
logs _and_ recorded as a condition on the `HelmRelease` resource. You can view
this condition by describing the `HelmRelease` resource using `kubectl`:

```console
kubectl describe -n <namespace> helmrelease/<name>
```

In case of a release upgrade failure, the error as returned by Helm will be
recorded in the message of `HelmUpgradeFailed`. If this does not give a
conclusive answer the logs will likely contain more information about what
happened during the release process:

```console
kubectl logs deploy/flux-helm-operator
```

## Manually performing a release to debug

When describing the `HelmRelease` and the logs did not give any clues, it may
help to perform the release manually using the same values as specified in the
`HelmRelease` resource. When no `.valuesFrom` are defined, this can be done
by making use of [`yq`](https://github.com/kislyuk/yq) (an extension to `jq`)
and `kubectl`:

```console
kubectl get helmrelease/<name> -n <namespace> -o yaml | yq .spec.values -y | helm upgrade -i <release name> -f - <chart>
```

## Getting Help

If you still have any questions about the Helm Operator:

- Invite yourself to the <a href="https://slack.cncf.io" target="_blank">CNCF community</a>
  slack and ask a question on the [#flux](https://cloud-native.slack.com/messages/flux/)
  channel.
- To be part of the conversation about Helm Operator's development, join the
  [flux-dev mailing list](https://lists.cncf.io/g/cncf-flux-dev).
- [File an issue.](https://github.com/fluxcd/helm-operator/issues/new)

Your feedback is always welcome!
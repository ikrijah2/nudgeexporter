---
title: Alerting rules
sort_rank: 3
---

# Alerting rules

Alerting rules allow you to define alert conditions based on Prometheus
expression language expressions and to send notifications about firing alerts
to an external service. Whenever the alert expression results in one or more
vector elements at a given point in time, the alert counts as active for these
elements' label sets.

Alerting rules are configured in Prometheus in the same way as [recording
rules](recording_rules.md).

### Defining alerting rules

Alerting rules are defined in the following syntax:

    ALERT <alert name>
      IF <expression>
      [ FOR <duration> ]
      [ LABELS <label set> ]
      [ ANNOTATIONS <label set> ]

The alert name must be a valid metric name.

The optional `FOR` clause causes Prometheus to wait for a certain duration
between first encountering a new expression output vector element (like an
instance with a high HTTP error rate) and counting an alert as firing for this
element. Elements that are active, but not firing yet, are in pending state.

The `LABELS` clause allows specifying a set of additional labels to be attached
to the alert. Any existing conflicting labels will be overwritten. The label
values can be templated.

The `ANNOTATIONS` clause specifies another set of labels that are not
identifying for an alert instance. They are used to store longer additional
information such as alert descriptions or runbook links. The annotation values
can be templated.

#### Templating

Label and annotation values can be templated using [console templates](https://prometheus.io/docs/visualization/consoles).
The `$labels` variable holds the label key/value pairs of an alert instance
and `$value` holds the evaluated value of an alert instance.

    # To insert a firing element's label values:
    {{ $labels.<labelname> }}
    # To insert the numeric expression value of the firing element:
    {{ $value }}

Examples:

    # Alert for any instance that is unreachable for >5 minutes.
    ALERT InstanceDown
      IF up == 0
      FOR 5m
      LABELS { severity = "page" }
      ANNOTATIONS {
        summary = "Instance {{ $labels.instance }} down",
        description = "{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 5 minutes.",
      }

    # Alert for any instance that have a median request latency >1s.
    ALERT APIHighRequestLatency
      IF api_http_request_latencies_second{quantile="0.5"} > 1
      FOR 1m
      ANNOTATIONS {
        summary = "High request latency on {{ $labels.instance }}",
        description = "{{ $labels.instance }} has a median request latency above 1s (current value: {{ $value }}s)",
      }

### Inspecting alerts during runtime

To manually inspect which alerts are active (pending or firing), navigate to
the "Alerts" tab of your Prometheus instance. This will show you the exact
label sets for which each defined alert is currently active.

For pending and firing alerts, Prometheus also stores synthetic time series of
the form `ALERTS{alertname="<alert name>", alertstate="pending|firing", <additional alert labels>}`.
The sample value is set to `1` as long as the alert is in the indicated active
(pending or firing) state, and a single `0` value gets written out when an alert
transitions from active to inactive state. Once inactive, the time series does
not get further updates.

### Sending alert notifications

Prometheus's alerting rules are good at figuring what is broken *right now*, but
they are not a fully-fledged notification solution. Another layer is needed to
add summarization, notification rate limiting, silencing and alert dependencies
on top of the simple alert definitions. In Prometheus's ecosystem, the
[Alertmanager](https://prometheus.io/docs/alerting/alertmanager/) takes on this
role. Thus, Prometheus may be configured to periodically send information about
alert states to an Alertmanager instance, which then takes care of dispatching
the right notifications. The Alertmanager instance may be configured via the
`-alertmanager.url` command line flag.

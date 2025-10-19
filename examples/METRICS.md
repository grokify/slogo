# SLO Metrics by Ontology Labels

This document shows the distribution of SLOs across different ontology label dimensions.


## Framework

| Directory                              | Red    | Use    | Custom | Total   |
| -------------------------------------- | ------ | ------ | ------ | ------- |
| ai-agents                              | 5      | -      | 15     | 20      |
| budgeting-method                       | -      | -      | 3      | 3       |
| red-metrics                            | 5      | -      | -      | 5       |
| saas-crm                               | -      | -      | 24     | 24      |
| treat-low-traffic-as-equally-important | -      | -      | 1      | 1       |
| use-metrics                            | -      | 11     | -      | 11      |

## Layer

| Directory                              | Service | Infrastructure | Business | Application | Total   |
| -------------------------------------- | ------- | -------------- | -------- | ----------- | ------- |
| ai-agents                              | 5       | -              | 15       | -           | 20      |
| budgeting-method                       | 3       | -              | -        | -           | 3       |
| red-metrics                            | 5       | -              | -        | -           | 5       |
| saas-crm                               | -       | -              | 24       | -           | 24      |
| treat-low-traffic-as-equally-important | 1       | -              | -        | -           | 1       |
| use-metrics                            | -       | 11             | -        | -           | 11      |

## Scope

| Directory                              | Customer Facing | Internal | Business Outcome | Total   |
| -------------------------------------- | --------------- | -------- | ---------------- | ------- |
| ai-agents                              | 8               | -        | 12               | 20      |
| budgeting-method                       | 3               | -        | -                | 3       |
| red-metrics                            | 5               | -        | -                | 5       |
| saas-crm                               | -               | -        | 24               | 24      |
| treat-low-traffic-as-equally-important | 1               | -        | -                | 1       |
| use-metrics                            | -               | 11       | -                | 11      |

## Audience

| Directory                              | Sre    | Engineering | Product | Executive | Customer Success | Total   |
| -------------------------------------- | ------ | ----------- | ------- | --------- | ---------------- | ------- |
| ai-agents                              | 5      | -           | 11      | 4         | -                | 20      |
| budgeting-method                       | 3      | -           | -       | -         | -                | 3       |
| red-metrics                            | 5      | -           | -       | -         | -                | 5       |
| saas-crm                               | -      | -           | 21      | 3         | -                | 24      |
| treat-low-traffic-as-equally-important | 1      | -           | -       | -         | -                | 1       |
| use-metrics                            | 11     | -           | -       | -         | -                | 11      |

## Category

| Directory                              | Availability | Latency | Throughput | Quality | Resource | Engagement | Conversion | Cost   | Security | Compliance | Total   |
| -------------------------------------- | ------------ | ------- | ---------- | ------- | -------- | ---------- | ---------- | ------ | -------- | ---------- | ------- |
| ai-agents                              | -            | 3       | -          | 9       | -        | 4          | -          | 4      | -        | -          | 20      |
| budgeting-method                       | 2            | 1       | -          | -       | -        | -          | -          | -      | -        | -          | 3       |
| red-metrics                            | 1            | 2       | 1          | 1       | -        | -          | -          | -      | -        | -          | 5       |
| saas-crm                               | -            | -       | -          | -       | -        | 12         | 12         | -      | -        | -          | 24      |
| treat-low-traffic-as-equally-important | 1            | -       | -          | -       | -        | -          | -          | -      | -        | -          | 1       |
| use-metrics                            | -            | -       | -          | -       | 11       | -          | -          | -      | -        | -          | 11      |

## Severity

| Directory                              | Critical | High   | Medium | Low    | Info   | Total   |
| -------------------------------------- | -------- | ------ | ------ | ------ | ------ | ------- |
| ai-agents                              | 6        | 10     | 4      | -      | -      | 20      |
| budgeting-method                       | 2        | 1      | -      | -      | -      | 3       |
| red-metrics                            | 2        | 2      | 1      | -      | -      | 5       |
| saas-crm                               | 5        | 14     | 5      | -      | -      | 24      |
| treat-low-traffic-as-equally-important | 1        | -      | -      | -      | -      | 1       |
| use-metrics                            | 4        | 5      | 2      | -      | -      | 11      |

## Tier

| Directory                              | P0     | P1     | P2     | P3     | Total   |
| -------------------------------------- | ------ | ------ | ------ | ------ | ------- |
| ai-agents                              | 11     | 5      | 4      | -      | 20      |
| budgeting-method                       | 2      | 1      | -      | -      | 3       |
| red-metrics                            | 3      | 2      | -      | -      | 5       |
| saas-crm                               | 9      | 10     | 5      | -      | 24      |
| treat-low-traffic-as-equally-important | 1      | -      | -      | -      | 1       |
| use-metrics                            | 6      | 5      | -      | -      | 11      |

## Metric Type

| Directory                              | Rate   | Errors | Duration | Availability | Utilization | Saturation | Satisfaction | Stickiness | Retention | Activation | Adoption | Efficiency | Total   |
| -------------------------------------- | ------ | ------ | -------- | ------------ | ----------- | ---------- | ------------ | ---------- | --------- | ---------- | -------- | ---------- | ------- |
| ai-agents                              | -      | 2      | 3        | -            | -           | -          | 3            | -          | -         | -          | -        | 8          | 16      |
| budgeting-method                       | -      | -      | 1        | 2            | -           | -          | -            | -          | -         | -          | -        | -          | 3       |
| red-metrics                            | 1      | 1      | 2        | 1            | -           | -          | -            | -          | -         | -          | -        | -          | 5       |
| saas-crm                               | -      | -      | -        | -            | -           | -          | -            | 3          | 5         | 3          | 5        | -          | 16      |
| treat-low-traffic-as-equally-important | -      | -      | -        | 1            | -           | -          | -            | -          | -         | -          | -        | -          | 1       |
| use-metrics                            | -      | 4      | -        | -            | 3           | 4          | -            | -          | -         | -          | -        | -          | 11      |

## Resource Type

| Directory                              | Cpu    | Memory | Disk   | Network | Gpu    | Total   |
| -------------------------------------- | ------ | ------ | ------ | ------- | ------ | ------- |
| ai-agents                              | -      | -      | -      | -       | -      | 0       |
| budgeting-method                       | -      | -      | -      | -       | -      | 0       |
| red-metrics                            | -      | -      | -      | -       | -      | 0       |
| saas-crm                               | -      | -      | -      | -       | -      | 0       |
| treat-low-traffic-as-equally-important | -      | -      | -      | -       | -      | 0       |
| use-metrics                            | 3      | 3      | 3      | 2       | -      | 11      |

## Domain

| Directory                              | Ai Ml  | Crm    | Saas   | Ecommerce | Fintech | Total   |
| -------------------------------------- | ------ | ------ | ------ | --------- | ------- | ------- |
| ai-agents                              | 20     | -      | -      | -         | -       | 20      |
| budgeting-method                       | -      | -      | -      | 2         | -       | 2       |
| red-metrics                            | -      | -      | -      | -         | -       | 0       |
| saas-crm                               | -      | 24     | -      | -         | -       | 24      |
| treat-low-traffic-as-equally-important | -      | -      | -      | -         | -       | 0       |
| use-metrics                            | -      | -      | -      | -         | -       | 0       |

## Journey Stage

| Directory                              | Acquisition | Activation | Engagement | Retention | Revenue | Referral | Total   |
| -------------------------------------- | ----------- | ---------- | ---------- | --------- | ------- | -------- | ------- |
| ai-agents                              | -           | -          | -          | -         | -       | -        | 0       |
| budgeting-method                       | -           | -          | -          | -         | -       | -        | 0       |
| red-metrics                            | -           | -          | -          | -         | -       | -        | 0       |
| saas-crm                               | -           | 3          | 8          | 5         | 3       | -        | 19      |
| treat-low-traffic-as-equally-important | -           | -          | -          | -         | -       | -        | 0       |
| use-metrics                            | -           | -          | -          | -         | -       | -        | 0       |

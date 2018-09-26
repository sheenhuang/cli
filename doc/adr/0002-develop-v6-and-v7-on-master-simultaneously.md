# 2. Develop v6 and v7 on master simultaneously

Date: 2018-09-25

## Status

Proposed

## Context

_The issue motivating this decision, and any context that influences or constrains the decision._

We have two teams developing the CLI. 
One (CLI) is building out functionality (such as a better UAA and login experience) for the current, stable v6 releases. 
One (VAT) is dropping compatibility with CAPI v2 to more quickly integrate with and build out CAPI v3.

For the last three weeks, 
VAT has been developing on a long-running _vat_ branch based on v6.39. 
Part of the goals for VAT include 
incorporating all 7-10 new epics that will be built on _master_,
along with sharing any bugfixes back and forth between the two branches.
VAT has been tracking which commits should be cherry-picked back and forth each day, 
resulting in significant overhead and bookkeeping.
Rather than continuously merging two long-running branches into each other, 
the VAT team proposes to develop on _master_ simultaneously with CLI.
This would allow the CAPI v3 acceleration work to proceed faster.

The CLI team is constrained by needing to continue shipping stable versions of v6. 
There would be increased risk of shipping unstable versions, 
given that the VAT team's work involves modifying existing code
for which there are no integration tests.
The VAT team would bear little of this risk
and would reap most of the benefit.

## Decision

The change that we're proposing or have agreed to implement.

## Consequences

What becomes easier or more difficult to do and any risks introduced by the change that will need to be mitigated.

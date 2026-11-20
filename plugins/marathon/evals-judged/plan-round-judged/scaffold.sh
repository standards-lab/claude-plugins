#!/usr/bin/env bash
# Builds plan-round-in-reply's fixture by running that case's scaffold, so both cases plan the
# same task. eval refuses a scaffold_script outside its case directory, so this one delegates.
exec bash "$(dirname "$0")/../../evals/plan-round-in-reply/scaffold.sh"

#!/bin/bash

mv /Users/noamfavier/Neoware/AutoSort AutoSort
mkdir -p AutoSort/config
mv /Users/noamfavier/Neoware/AutoSort/config/default_rules.yaml AutoSort/config/
mkdir -p AutoSort/internal
mkdir -p AutoSort/internal/fileops
mkdir -p AutoSort/internal/llm
mkdir -p AutoSort/internal/rules
mkdir -p AutoSort/internal/sorter
mv /Users/noamfavier/Neoware/AutoSort/cmd/AutoSort/main.go AutoSort/cmd/AutoSort/
mv /Users/noamfavier/Neoware/AutoSort/internal/fileops/*.go AutoSort/internal/fileops/
mv /Users/noamfavier/Neoware/AutoSort/internal/llm/*.go AutoSort/internal/llm/
mv /Users/noamfavier/Neoware/AutoSort/internal/rules/*.go AutoSort/internal/rules/
mv /Users/noamfavier/Neoware/AutoSort/internal/sorter/*.go AutoSort/internal/sorter/
mkdir -p AutoSort/pkg
mv /Users/noamfavier/Neoware/AutoSort/README.md AutoSort/
mv /Users/noamfavier/Neoware/AutoSort/autosort_suggestion.md AutoSort/
```
# projman

Personal project management tool.

## Installation

This is very much a prject for personal benefit, so many things are tailored towards me personally currently. If more interest is show I may make changes to make it more generic.

```sh
git clone https://github.com/vcokltfre/projman.git
cd projman
./build
```

## Commands

```sh
proj start [name]   - start a new project
proj close [name]   - close a project
proj restore [name] - restore a project directory (does not re-start the project)
proj list           - list all projects
proj cleanup        - cleanup stale projects
proj cleanup-all    - cleanup stale and unregistered projects
proj validate       - validate that all registered projects exist
```

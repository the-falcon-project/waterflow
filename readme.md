# Waterflow

Waterflow is a lightweight task scheduler written in Go that allows you to define and automate workflows through YAML configuration files. It simplifies the scheduling and execution of tasks, making it easier to manage complex workflows.

## Table of Contents

- [Getting Started](#getting-started)
- [Usage](#usage)
- [YAML Structure](#yaml-structure)
- [Roadmap](#roadmap)
- [Contributing](#contributing)

## Getting Started

### Installation

To get started with Waterflow, you'll need to [download the latest release](https://github.com/the-falcon-project/waterflow) or clone the repository and build it locally.

```bash
# Clone the repository
git clone https://github.com/the-falcon-project/waterflow

# Build the project
cd waterflow
go build -o waterflow
```
## Configuration
Create a YAML configuration file to define your workflows and tasks. Refer to the YAML Structure section for details on how to structure your configuration.

## Usage
Waterflow provides a command-line interface (CLI) for managing workflows and tasks. Here are some common commands:

```bash
# Create a new workflow
waterflow create workflow my-workflow.yaml

# List all workflows
waterflow list workflows

# Run a specific workflow
waterflow run my-workflow-id

# Check the status of a workflow
waterflow status my-workflow-id

# Stop a running workflow
waterflow stop my-workflow-id

# Delete a workflow from the scheduler
waterflow delete my-workflow-id 

```
Check out how to schedule a flow with Waterflow [here](https://github.com/the-falcon-project/waterflow/wiki/Waterflow-Scheduling).

Refer to the YAML Structure Guide for a detailed explanation of the configuration format.

## Roadmap
Our project roadmap outlines upcoming features and improvements. Feel free to check out the Roadmap to see what's planned and provide your feedback.

> Will add this in some time.

## Contributing
We welcome contributions from the community! If you'd like to contribute to Waterflow, please follow our [Contribution Guidelines](contribution.md) and open a pull request.


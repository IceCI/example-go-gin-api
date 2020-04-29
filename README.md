# IceCI Go example - Gin API 

This repository is a small example to help you get started building **IceCI** pipelines for Go applications. 

The application itself is a very simplified example of a web API with tests. It's not by any means a guideline on building Go applications - its only purpose is showcasing how to build IceCI pipelines for Go applications.

This repository is a *GitHub template repository*, so please feel free to create new repositories based on it and mess around with the code and the pipeline config. Please also check the information below for a list of prerequisites needed to run the pipeline in IceCI.

# Setting up IceCI


To launch the pipeline in IceCI you have to create 2 secrets - both of which are explicitly referenced in the `.iceci.yaml` file. 

* `dockerhub` - docker hub credentials with `docker.io` set as registry.

* `slack-webhook` - a generic secret with hook for Slack notifications - it can be ignored by commenting out the `slack-notify` step as well as the `slack-notify-error` failure handler in the pipeline definition file. 

You can also find some additional info in the comments of the `.iceci.yaml` file itself.


---

Kept cool &#x1f9ca; by [Icetek](https://icetek.io/)

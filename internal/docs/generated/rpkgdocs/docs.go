// Code generated by "mdtogo"; DO NOT EDIT.
package rpkgdocs

var RpkgShort = `Manage packages.`
var RpkgLong = `
The ` + "`" + `rpkg` + "`" + ` command group contains subcommands for managing packages and revisions.
`

var ApproveShort = `Approve a proposal to publish a package revision.`
var ApproveLong = `
  kpt alpha rpkg approve [PACKAGE_REV_NAME...] [flags]

Args:

  PACKAGE_REV_NAME...:
    The name of one or more package revisions. If more than
    one is provided, they must be space-separated.
`
var ApproveExamples = `
  # approve package revision blueprint-91817620282c133138177d16c981cf35f0083cad
  $ kpt alpha rpkg approve blueprint-91817620282c133138177d16c981cf35f0083cad --namespace=default
`

var CloneShort = `Create a clone of an existing package revision.`
var CloneLong = `
  kpt alpha rpkg clone SOURCE_PACKAGE_REV TARGET_PACKAGE_NAME [flags]

Args:

  SOURCE_PACKAGE_REV:
    The source package that will be cloned to create the new package revision.
    The types of sources are supported:
  
      * OCI: A URI to a OCI repository must be provided. 
        oci://oci-repository/package-name
      * Git: A URI to a git repository must be provided.
        https://git-repository.git/package-name
      * Package: The name of a package revision already available in the
        repository.
        blueprint-e982b2196b35a4f5e81e92f49a430fe463aa9f1a
  
  TARGET_PACKAGE_NAME:
    The name of the new package.
  

Flags:

  --directory
    Directory within the repository where the upstream
    package revision is located. This only applies if the source package is in git
    or oci.
  
  --ref
    Ref in the repository where the upstream package revision
    is located (branch, tag, SHA). This only applies when the source package
    is in git.
  
  --repository
    Repository to which package revision will be cloned
    (downstream repository).
  
  --revision
    Revision for the new package.
  
  --strategy
    Update strategy that should be used when updating the new
    package revision. Must be one of: resource-merge, fast-forward,  or 
    force-delete-replace. The default value is resource-merge.
`
var CloneExamples = `
  # clone the blueprint-e982b2196b35a4f5e81e92f49a430fe463aa9f1a package and create a new package revision called
  # foo in the blueprint repository and set it at revision v1.
  $ kpt alpha rpkg clone blueprint-e982b2196b35a4f5e81e92f49a430fe463aa9f1a foo --repository blueprint --revision v1

  # clone the git repository at https://github.com/repo/blueprint.git at reference base/v0 and in directory base. The new
  # package revision will be created in repository blueprint and namespace default.
  $ kpt alpha rpkg clone https://github.com/repo/blueprint.git bar --repository=blueprint --ref=base/v0 --namespace=default --directory=base
`

var CopyShort = `Create a new package revision from an existing one.`
var CopyLong = `
  kpt alpha rpkg copy SOURCE_PACKAGE_REV_NAME TARGET_PACKAGE_NAME [flags]

Args:

  SOURCE_PACKAGE_REV_NAME:
    The name of the package revision that will be used as the source
    for creating a new package revision.
  
  TARGET_PACKAGE_NAME:
    The name of the new package.

Flags:

  --repository
    Repository in which the new package revision will be created.
  
  --revision
    Revision for the new package.
`
var CopyExamples = `
  # create a new package foo from package blueprint-b47eadc99f3c525571d3834cc61b974453bc6be2 
  $ kpt alpha rpkg copy blueprint-b47eadc99f3c525571d3834cc61b974453bc6be2 foo --repository=blueprint --revision=v10 --namespace=default
`

var DelShort = `Delete a package revision.`
var DelLong = `
  kpt alpha rpkg del PACKAGE_REV_NAME... [flags]

Args:

  PACKAGE_REV_NAME...:
    The name of one or more package revisions. If more than
    one is provided, they must be space-separated.
`
var DelExamples = `
  # remove package revision blueprint-e982b2196b35a4f5e81e92f49a430fe463aa9f1a from the default namespace
  $ kpt alpha rpkg del blueprint-e982b2196b35a4f5e81e92f49a430fe463aa9f1a --namespace=default
`

var GetShort = `List package revisions in registered repositories.`
var GetLong = `
  kpt alpha rpkg get [PACKAGE_REV_NAME] [flags]

Args:

  PACKAGE_REV_NAME:
    The name of a package revision. If provided, only that specific
    package revision will be shown. Defaults to showing all package
    revisions from all repositories.

Flags:

  --name
    Name of the packages to get. Any package whose name contains 
    this value will be included in the results.
  
  --revision
    Revision of the package to get. Any package whose revision
    matches this value will be included in the results.
`
var GetExamples = `
  # get a specific package revision in the default namespace
  $ kpt alpha rpkg get blueprint-e982b2196b35a4f5e81e92f49a430fe463aa9f1a --namespace=default

  # get all package revisions in the bar namespace
  $ kpt alpha rpkg get --namespace=bar

  # get all package revisions with revision v0
  $ kpt alpha rpkg get --revision=v0
`

var InitShort = `Initializes a new package in a repository.`
var InitLong = `
  kpt alpha rpkg init PACKAGE_NAME [flags]

Args:

  PACKAGE_NAME:
    The name of the new package.

Flags:

  --repository
    Repository in which the new package will be created.
  
  --revision
    Revision of the new package. The default value if v1.
  
  --description
    short description of the package
  
  --keywords
    list of keywords for the package
  
  --site
    link to page with information about the package
`
var InitExamples = `
  # create a new package named foo in the repository blueprint.
  $ kpt alpha rpkg init foo --namespace=default --repository=blueprint
`

var ProposeShort = `Propose that a package revision should be published.`
var ProposeLong = `
  kpt alpha rpkg propose [PACKAGE_REV_NAME...] [flags]

Args:

  PACKAGE_REV_NAME...:
    The name of one or more package revisions. If more than
    one is provided, they must be space-separated.
`
var ProposeExamples = `
  # propose that package revision blueprint-91817620282c133138177d16c981cf35f0083cad should be finalized.
  $ kpt alpha rpkg propose blueprint-91817620282c133138177d16c981cf35f0083cad --namespace=default
`

var PullShort = `Pull the content of the package revision.`
var PullLong = `
  kpt alpha rpkg pull PACKAGE_REV_NAME [DIR] [flags]

Args:

  PACKAGE_REV_NAME:
    The name of a an existing package revision in a repository.
  
  DIR:
    A local directory where the package manifests will be written.
    If not provided, the manifests are written to stdout.
`
var PullExamples = `
  # pull the content of package revision blueprint-d5b944d27035efba53836562726fb96e51758d97
  $ kpt alpha rpkg pull blueprint-d5b944d27035efba53836562726fb96e51758d97 --namespace=default
`

var PushShort = `Push resources to a package revision.`
var PushLong = `
  kpt alpha rpkg push PACKAGE_REV_NAME [DIR] [flags]

Args:

  PACKAGE_REV_NAME:
    The name of a an existing package revision in a repository.
  
  DIR:
    A local directory with the new manifest. If not provided,
    the manifests will be read from stdin.
`
var PushExamples = `
  # update the package revision blueprint-f977350dff904fa677100b087a5bd989106d0456 with the resources
  # in the ./package directory
  $ kpt alpha rpkg push blueprint-f977350dff904fa677100b087a5bd989106d0456 ./package --namespace=default
`

var RejectShort = `Reject a proposal to publish a package revision.`
var RejectLong = `
  kpt alpha rpkg reject [PACKAGE_REV_NAME...] [flags]

Args:

  PACKAGE_REV_NAME...:
    The name of one or more package revisions. If more than
    one is provided, they must be space-separated.
`
var RejectExamples = `
  # reject the proposal for package revision blueprint-8f9a0c7bf29eb2cbac9476319cd1ad2e897be4f9
  $ kpt alpha rpkg reject blueprint-8f9a0c7bf29eb2cbac9476319cd1ad2e897be4f9 --namespace=default
`

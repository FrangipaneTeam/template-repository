<div align="center">
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/releases/latest">
      <img alt="Latest release" src="https://img.shields.io/github/v/release/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=starship&color=C9CBFF&logoColor=D9E0EE&labelColor=302D41&include_prerelease&sort=semver" />
    </a>
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/pulse">
      <img alt="Last commit" src="https://img.shields.io/github/last-commit/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=starship&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
    </a>
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/stargazers">
      <img alt="Stars" src="https://img.shields.io/github/stars/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=starship&color=c69ff5&logoColor=D9E0EE&labelColor=302D41" />
    </a>
    <a href="https://github.com/FrangipaneTeam/terraform-plugin-framework-supertype/issues">
      <img alt="Issues" src="https://img.shields.io/github/issues/FrangipaneTeam/terraform-plugin-framework-supertype?style=for-the-badge&logo=bilibili&color=F5E0DC&logoColor=D9E0EE&labelColor=302D41" />
    </a>
</div>

# terraform-plugin-framework-supertype

supertype is a custom type of Terraform type for resources and datasources, along with a common field that enables you to manipulate Go object and Terrform object. supertype is compatible with [tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs).

This is a try to solve these issues :

* Don't repeat yourself : common fields applied on resources and datasources.
* Auto format attributes markdown description with validators and plan modifiers descriptions, default values...

## Documentation

For more information about the supertype, please refer to the [documentation](https://github.frangipane.io/terraform/supertype/why).
"""
Output changelog for each sub project
"""

def _version_git_chglog(ctx):
    for i in ["nada-sensio", "nada-serve", "nada-transform"]:
        output_file = ctx.actions.declare_file(i + "/CHANGELOG.md")
        ctx.actions.run(output_file, executable = "git-chglog", )
    pass

version = rule(
    implementation = _version_git_chglog,
)

# Caching build results

By-default Anzer stores build cache in the `~/.anzer_cache/` folder (for Linux and OS X). 
The folder contains zip files named with "package-function-commit hash" that are ready to deploy to the platform.

The actual versions (git commit hashes) of the functions stored in the `anzer.sum` file. `anzer.sum` can be modified
manually or with any csv editor or with `anzer sum update -f repository/funtion/name` command.
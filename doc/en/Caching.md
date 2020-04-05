# Caching build results

By-default anzer stores build cache in `~/.anzer_cache/`. 
The folder contains zip files named with "package-function-commit hash" and ready to deploy to the platform.

Actual versions (git commit hashes) of the functions are stored in the `anzer.sum` file.
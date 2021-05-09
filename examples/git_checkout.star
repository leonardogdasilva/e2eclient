

# git stash; \
# git checkout $(2); \
# git fetch --all; \
# git reset --hard origin/$(2); \

git_stash = run_local("git stash")
print(git_stash)
git_checkout = run_local("git checkout " + args.target)
print(git_checkout)

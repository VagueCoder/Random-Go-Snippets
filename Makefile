#!make

# Declared variables
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
REPO_ROOT = ${shell go env GOPATH}/src/github.com/VagueCoder/Random-Go-Snippets
ENV_FILE=env.sh

##################################### MAINTENANCE ####################################

# Include from file and export variables so the rules can make use
include ${ENV_FILE}
export $(shell sed 's/=.*//' ${ENV_FILE})

# Commit all the changes in local repo, in current branch
commit:
	- git add ${REPO_ROOT}
ifdef c
	- git commit -m "${c}"
else
	- git commit -m "Corrections"
endif

# Commit and push the changes to remote git repo
push:
	- make commit c="${c}"
	- git push -u origin ${GIT_BRANCH}

pushall:
	- make commit c="${c}"
	- git push -u origin --all

# Merge all changes of current branch with specified branch
merge:
ifeq (,$(and $(filter Changes not staged for commit, $(shell git status)), $(filter Changes to be committed, $(shell git status))))
	- @echo "Changes ommitted/up-to-date for current working branch. Proceeding...";
ifdef to
ifeq (,$(filter $(to), $(GIT_BRANCH)))
ifneq (,$(filter $(to), $(shell git branch)))
	- @echo "Branch '${to}' found. Proceeding...";
	- $(eval CURRENT_BRANCH := $(GIT_BRANCH))
	- @git checkout ${to};
	- @git merge $(CURRENT_BRANCH);
	- @git checkout $(CURRENT_BRANCH);
	- @echo "All changes of '$(CURRENT_BRANCH)' merged with '$(to)'. Back to '$(CURRENT_BRANCH)'.";
else
	- @echo "Exited. Branch '${to}' not found.";
	- @exit 0;
endif
else
	- @echo "Exited. Current branch and merge-to branch cannot be same.";
	- @exit 0;
endif
else
	- @echo "Exited. Provide merge-to branch as to=<branch_name> and retry.";
	- @exit 0;
endif
else
	- @echo "Exited. Please do the add/rm/commit in current branch and retry.";
	- @exit 0;
endif

######################################################################################

#################################### Fast-Fibonacci ##################################

run_fast_fibonacci:
	@go build -o=${GOBIN}/fast_fibonacci ${REPO_ROOT}/Fast-Fibonacci;
	${GOBIN}/fast_fibonacci;

test_fast_fibonacci:
	go test -v ${REPO_ROOT}/Fast-Fibonacci/...

######################################################################################
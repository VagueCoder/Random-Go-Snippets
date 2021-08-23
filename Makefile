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

#################################### Test-Cases ##################################

test_fast_fibonacci:
	go test -v ${REPO_ROOT}/Fast-Fibonacci/...

test_files_filtered:
	go test -v ${REPO_ROOT}/Files-Filtered-By-Time-Window/...

test_formatted_time_marshalling:
	go test -v ${REPO_ROOT}/Formatted-Time-Marshalling/...

test_random_strings:
	go test -v ${REPO_ROOT}/Random-Strings/...

test_hashmap:
	go test -v ${REPO_ROOT}/Go-HashMap/...

test_broadcaster:
	go test -v ${REPO_ROOT}/Message-Broadcaster/...

test_find_gcd:
	go test -v ${REPO_ROOT}/Find-GCD/...

######################################################################################

######################################## General #####################################

clean_cache:
	@go clean -cache ./... && echo "Bye-bye Cache!"

######################################################################################
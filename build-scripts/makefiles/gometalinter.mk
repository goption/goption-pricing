# Copyright (C) 2018 Brian Hazeltine <onwsk8r@gmail.com> https://wasthat.me

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published
# by the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

## Gometalinter
# https://github.com/alecthomas/gometalinter

CURL ?= curl -sfL# The command to get files
LINTER_OPTIONS ?= --deadline 599s# Arguments to golangci-lint
LINTER_BINARY ?= gometalinter# Name of the binary of this linter

_pre_lint:
ifeq (,$(shell command $(LINTER_BINARY)))
	$(CURL) https://git.io/vp6lP | sh
endif
.PHONY: _install_lint

_lint:
	$(LINTER_BINARY) $(LINTER_OPTIONS) ./...
.PHONY: _lint

_lint_changed: lint
	@echo "_lint_changed is not implemented. Fell back to _lint."
.PHONY: _lint_changed

_clean_lint:
	@echo "_clean_lint: Nothing to do."
.PHONY: _clean_lint

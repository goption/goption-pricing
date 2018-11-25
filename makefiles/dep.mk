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

# Dep
# https://golang.github.io/dep/

OS ?= $(shell uname -s)
CURL ?= curl -sfL# The command to get files

DEP_BINARY ?= dep# The name of the dep binary
VENDOR_DIR ?= vendor# The directory where vendored dependencies live

_pre_depend:
ifeq (,$(shell command -v $(DEP_BINARY)))
ifeq ($(OS), Darwin)
	brew install dep
else
	$(CURL) https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
endif
.PHONY: _pre_depend

_depend:
	$(DEP_BINARY) ensure
.PHONY: _depend

_update_depend:
	@echo "_update_depend: This command has not been implemented yet."
.PHONY: _update_depend

_clean_depend:
	rm -Rf $(VENDOR_DIR)
.PHONY: _clean_depend

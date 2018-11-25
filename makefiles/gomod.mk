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

# Golang module support using good ol `go mod`

VENDOR_DIR ?= vendor# The directory where vendored dependencies live

# Using regular go mod
_pre_depend:
	@echo "_pre_depend: nothing to do"
.PHONY: _pre_depend

_depend: # Fetch modules with `go mod download`
	go mod download
.PHONY: _depend

_update_depend: # Update (I think) modules with `go mod tidy`
	go mod tidy
.PHONY: _update_depend

_clean_depend:
	@echo "_clean_depend: this command has not been implemented"
.PHONY: _clean_depend

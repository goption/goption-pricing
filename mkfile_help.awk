#!/usr/bin/env awk
# mkfile_help.awk: Autogenerate help from Makefiles
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

BEGIN {
    # Number of columns in the terminal
    # Needed when we start worrying about line wrapping
    # "tput cols" | getline term_cols

    # Initialize the variables we use (all assoc arrays)
    funcs[""] = 0
    vars[""] = 0
    var_type[""] = 0
    var_valu[""] = 0
    mf_vars[""] = 0

    # Colorize the output so even the marketing team loves it
    # https://misc.flogisoft.com/bash/tip_colors_and_formatting
    uline = "\033[1;4m"
    red = "\033[1;31m"
    green = "\033[1;32m"
    yellow = "\033[1;33m"
    cyan = "\033[1;36m"
    reset = "\033[0m"

    print green "*** Only listing functions and variables with comments ***\n" reset
}

# Print the filename with no path at the beginning of each file
BEGINFILE { print uline red gensub(/.*\//, "", 1, FILENAME) reset }

# The main block only parses the job and variable definitions into
# arrays that can be manipulated later. While this is npt strictly
# necessary for this use case, it is a nice separation of concerns.
{
    # Match functions, this is a solid regex
    if (match($0, "^([^:]+):[^#]+#+(.+)$", row)) {
        funcs[row[1]] = row[2]
    # Match variables, this is a terrible regex
    } else if (match($0, "^([A-Za-z0-9_-]+) *([?:+]:?)= *([^#]*)#+(.+)", row)) {
        # These if statements are to keep variables that are defined in Makefile
        # from being printed again if they are redefined in subsequent *.mk files
        # with `?=` as they are inherently overridden by the definition in the
        # main Makefile. In order for this to work, the Makefiles must be passed
        # to awk in the order of precedence, which they are if using $(MAKEFILE_LIST)
        if (mf_vars[row[1]] < 1 && row[3] != "?") {
            vars[row[1]] = row[4]
            var_type[row[1]] = row[2]
            var_valu[row[1]] = row[3]
        }
        if (FILENAME == "Makefile") {
            mf_vars[row[1]]++
        }

    }
}

# The endfile block contains all of the display logic. I used 0 and 1 to
# represent booleans because it looks like there are no actual booleans.
ENDFILE {
    col1_width = 20
    var_format = cyan "%-" col1_width "s" reset "%s (" cyan "%s=%s" reset ")\n"
    func_format = yellow "%-" col1_width "s" reset "%s\n"

    has_var = 0
    has_func = 0

    # Variables...
    for (i in vars) {
        if (i != "") {
            # Unfortunately long variables wrap around and don't indent
            # themselves. It looks like this problem could be fixed with
            # the solution in this SE:
            # https://unix.stackexchange.com/questions/280199/wrap-and-indent-text-using-coreutils
            printf(var_format, i, vars[i], var_type[i], var_valu[i])
            has_var = 1
        }
    }
    delete(vars)
    delete(var_type)
    delete(var_valu)

    if (has_var == 0) {
        print cyan "This file does not define any new variables" reset
    }

    # Functions...
    for (i in funcs) {
        if (i != "") {
            printf(func_format, i "()", funcs[i])
            has_func = 1
        }
    }
    delete(funcs)

    if (has_func == 0) {
        print yellow "This file does contain any user functions" reset
    }

    print "" # Newline after each file
}

END {
    print green "Generated on " strftime("%a %FT%T%z") reset
}

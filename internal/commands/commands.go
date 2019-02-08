package commands

import (
	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pakeos/internal/commands/copy"
	"github.com/pake-go/pakeos/internal/commands/install"
	"github.com/pake-go/pakeos/internal/commands/link"
	"github.com/pake-go/pakeos/internal/commands/move"
	"github.com/pake-go/pakeos/internal/commands/remove"
	"github.com/pake-go/pakeos/internal/commands/run"
	"github.com/pake-go/pakeos/internal/commands/set"
	"github.com/pake-go/pakeos/internal/commands/setdefault"
	"github.com/pake-go/pakeos/internal/commands/uninstall"
	"github.com/pake-go/pakeos/internal/commands/unlink"
)

// Candidates is a list of commands that might be used in the source file.
var Candidates = []pakelib.CommandCandidate{
	copyCandidate,
	installCandidate,
	linkCandidate,
	moveCandidate,
	removeCandidate,
	runCandidate,
	setCandidate,
	setdefaultCandidate,
	uninstallCandidate,
	unlinkCandidate,
}

var copyCandidate = pakelib.CommandCandidate{
	Validator:   &copy.CopyValidator{},
	Constructor: copy.New,
}

var installCandidate = pakelib.CommandCandidate{
	Validator:   &install.InstallValidator{},
	Constructor: install.New,
}

var linkCandidate = pakelib.CommandCandidate{
	Validator:   &link.LinkValidator{},
	Constructor: link.New,
}

var moveCandidate = pakelib.CommandCandidate{
	Validator:   &move.MoveValidator{},
	Constructor: move.New,
}

var removeCandidate = pakelib.CommandCandidate{
	Validator:   &remove.RemoveValidator{},
	Constructor: remove.New,
}

var runCandidate = pakelib.CommandCandidate{
	Validator:   &run.RunValidator{},
	Constructor: run.New,
}

var setCandidate = pakelib.CommandCandidate{
	Validator:   &set.SetValidator{},
	Constructor: set.New,
}

var setdefaultCandidate = pakelib.CommandCandidate{
	Validator:   &setdefault.SetDefaultValidator{},
	Constructor: setdefault.New,
}

var uninstallCandidate = pakelib.CommandCandidate{
	Validator:   &uninstall.UninstallValidator{},
	Constructor: uninstall.New,
}

var unlinkCandidate = pakelib.CommandCandidate{
	Validator:   &unlink.UnlinkValidator{},
	Constructor: unlink.New,
}

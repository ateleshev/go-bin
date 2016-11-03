package logger

type msgWriter func(Logger, *msg)

func getMsgWriter(mode Mode) msgWriter { // {{{
	switch mode {
	case ModeInfo:
		return msgInfoWriter
	case ModeError:
		return msgErrorWriter
	case ModeDebug:
		return msgDebugWriter
	case ModeAccess:
		return msgAccessWriter
	}

	panic("Unknown mode")

	return nil
} // }}}

func msgInfoWriter(logger Logger, m *msg) { // {{{
	if logger != nil {
		if m.Formatted() {
			logger.Infof(m.format, m.values...)
		} else {
			logger.Info(m.values...)
		}
	} else {
		println("msgInfoWriter logger - nil")
	}
} // }}}

func msgErrorWriter(logger Logger, m *msg) { // {{{
	if logger != nil {
		if m.Formatted() {
			logger.Errorf(m.format, m.values...)
		} else {
			logger.Error(m.values...)
		}
	} else {
		println("msgInfoWriter logger - nil")
	}
} // }}}

func msgDebugWriter(logger Logger, m *msg) { // {{{
	if logger != nil {
		if m.Formatted() {
			logger.Debugf(m.format, m.values...)
		} else {
			logger.Debug(m.values...)
		}
	} else {
		println("msgInfoWriter logger - nil")
	}
} // }}}

func msgAccessWriter(logger Logger, m *msg) { // {{{
	if logger != nil {
		if m.Formatted() {
			logger.Accessf(m.format, m.values...)
		} else {
			logger.Access(m.values...)
		}
	} else {
		println("msgInfoWriter logger - nil")
	}
} // }}}

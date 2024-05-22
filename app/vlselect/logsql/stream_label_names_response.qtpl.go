// Code generated by qtc from "stream_label_names_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// StreamLabelNamesResponse formats /select/logsql/stream_label_names response

//line app/vlselect/logsql/stream_label_names_response.qtpl:4
package logsql

//line app/vlselect/logsql/stream_label_names_response.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vlselect/logsql/stream_label_names_response.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vlselect/logsql/stream_label_names_response.qtpl:4
func StreamStreamLabelNamesResponse(qw422016 *qt422016.Writer, names []string) {
//line app/vlselect/logsql/stream_label_names_response.qtpl:4
	qw422016.N().S(`{"names":[`)
//line app/vlselect/logsql/stream_label_names_response.qtpl:7
	if len(names) > 0 {
//line app/vlselect/logsql/stream_label_names_response.qtpl:8
		qw422016.N().Q(names[0])
//line app/vlselect/logsql/stream_label_names_response.qtpl:9
		for _, v := range names[1:] {
//line app/vlselect/logsql/stream_label_names_response.qtpl:9
			qw422016.N().S(`,`)
//line app/vlselect/logsql/stream_label_names_response.qtpl:10
			qw422016.N().Q(v)
//line app/vlselect/logsql/stream_label_names_response.qtpl:11
		}
//line app/vlselect/logsql/stream_label_names_response.qtpl:12
	}
//line app/vlselect/logsql/stream_label_names_response.qtpl:12
	qw422016.N().S(`]}`)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
}

//line app/vlselect/logsql/stream_label_names_response.qtpl:15
func WriteStreamLabelNamesResponse(qq422016 qtio422016.Writer, names []string) {
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	StreamStreamLabelNamesResponse(qw422016, names)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
}

//line app/vlselect/logsql/stream_label_names_response.qtpl:15
func StreamLabelNamesResponse(names []string) string {
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	WriteStreamLabelNamesResponse(qb422016, names)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	qs422016 := string(qb422016.B)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
	return qs422016
//line app/vlselect/logsql/stream_label_names_response.qtpl:15
}

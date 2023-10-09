// Code generated by qtc from "transfer_list.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:2
package templates

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:2
import "time"

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:3
import "github.com/sollniss/kakeibo/domain"

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:4
import "github.com/sollniss/kakeibo/usecases/transfer"

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:7
type TransferList struct {
	Base
	Header
	Footer

	Transfers []domain.Transfer
	Current   string
	Data      transfer.ViewListData
}

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:18
func (t *TransferList) StreamBody(qw422016 *qt422016.Writer) {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:19
	t.StreamHeader(qw422016, t.Current, t.Data.Month)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:19
	qw422016.N().S(`<div class="content">`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:22
	start := time.Time{}
	current := start

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:26
	for _, trans := range t.Transfers {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:27
		if !trans.Date.Equal(current) {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:27
			qw422016.N().S(`<div class="list__heading">`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:29
			qw422016.E().S(japaneseDays.Replace(trans.Date.Format("1月2日（Mon）")))
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:29
			qw422016.N().S(`<span class="list__details">`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:31
			if t.Data.DaySums[trans.Date.Unix()] >= 0 {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:32
				qw422016.N().DL(t.Data.DaySums[trans.Date.Unix()])
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:33
			} else {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:34
				qw422016.N().DL(-t.Data.DaySums[trans.Date.Unix()])
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:35
			}
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:35
			qw422016.N().S(`</span></div>`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:40
			current = trans.Date

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:42
		}
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:42
		qw422016.N().S(`<div class="list" id="`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:44
		qw422016.E().S(trans.Date.Format("2006-01-02"))
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:44
		qw422016.N().S(`"><label class="list__elem accordion" for="`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.N().S(`-`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.N().DUL(uint64(trans.ID))
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.N().S(`-item" id="`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.N().S(`-`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.N().DUL(uint64(trans.ID))
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:45
		qw422016.N().S(`"><input class="accordion__switch" id="`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:46
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:46
		qw422016.N().S(`-`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:46
		qw422016.N().DUL(uint64(trans.ID))
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:46
		qw422016.N().S(`-item" type="checkbox"><div class="accordion__col1"><div class="cat-icon`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:48
		qw422016.N().S(` `)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:48
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:48
		qw422016.N().S(`-color12"><img class="cat-icon__img" src="/static/img/icons/`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:49
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:49
		qw422016.N().S(`-12.svg" alt="食費"></div></div><div class="accordion__col2">`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:53
		if trans.Amount >= 0 {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:54
			qw422016.N().DL(trans.Amount)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:55
		} else {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:56
			qw422016.N().DL(-trans.Amount)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:57
		}
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:57
		qw422016.N().S(`</div><div class="accordion__col3 accordion__content--mobile">`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:59
		qw422016.E().S(trans.Comment)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:59
		qw422016.N().S(`</div><div class="accordion__col4"><span class="person person--color2">D</span></div><div class="accordion__content"><a class="modal__open accordion__btn" data-modal="edit" data-target="/`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:64
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:64
		qw422016.N().S(`/2970" data-date="2023-10-02" data-amount="1181" data-type="12" data-person="2" data-comment="ランチ、シリアル、ライフ">編集</a><a class="modal__open accordion__btn" data-modal="delete" data-target="/`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:65
		qw422016.E().S(t.Current)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:65
		qw422016.N().S(`/2970">削除</a></div></label></div>`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:69
	}
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:69
	qw422016.N().S(`</div>`)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:71
	t.StreamFooter(qw422016, t.Current, t.Data.Month)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
}

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
func (t *TransferList) WriteBody(qq422016 qtio422016.Writer) {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	t.StreamBody(qw422016)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
}

//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
func (t *TransferList) Body() string {
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	t.WriteBody(qb422016)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	qs422016 := string(qb422016.B)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
	return qs422016
//line transport/http/html/quicktemplate/templates/transfer_list.qtpl:72
}

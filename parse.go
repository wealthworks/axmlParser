package axmlParser

import (
	"archive/zip"
	"io"
	"io/ioutil"
)

func ParseApk(apkpath string, listener Listener) (*Parser, error) {
	r, err := zip.OpenReader(apkpath)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return parseApkR(&r.Reader, listener)
}

func ParseApkReader(rd io.ReaderAt, size int64, listener Listener) (*Parser, error) {
	r, err := zip.NewReader(rd, size)
	if err != nil {
		return nil, err
	}
	return parseApkR(r, listener)
}

func parseApkR(r *zip.Reader, listener Listener) (parser *Parser, err error) {

	var xmlf *zip.File

	for _, f := range r.File {
		if f.Name != "AndroidManifest.xml" {
			continue
		}
		xmlf = f
		break
	}

	if xmlf == nil {
		return nil, err
	}

	rc, err := xmlf.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	bs, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	parser = New(listener)
	err = parser.Parse(bs)
	if err != nil {
		return nil, err
	}
	return parser, nil
}

func ParseAxml(axmlpath string, listener Listener) (*Parser, error) {
	bs, err := ioutil.ReadFile(axmlpath)
	if err != nil {
		return nil, err
	}
	parser := New(listener)
	err = parser.Parse(bs)
	if err != nil {
		return nil, err
	}
	return parser, nil
}

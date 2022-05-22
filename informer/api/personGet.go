package api

type ReqPersonGet struct {

}

type RespPersonGet struct {
	person string
}

func (req *ReqPersonGet) Execute() ( resp *RespPersonGet, err error) {
	out := RespPersonGet{
		person: "KEKbl4",
	}

	return &out, nil
}
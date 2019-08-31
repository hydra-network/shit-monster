package main

import (
	"./utils"
	"github.com/ajstarks/svgo"
	"log"
	"net/http"
)

type Hip struct {
	height int
	power int
}

type Shin struct {
	height int
	power int
}

type Joint struct {
	degree int
}

type Knee struct {
    id string
	hip Hip
    shin Shin
    joint Joint
}

func main() {
	http.Handle("/knee", http.HandlerFunc(drawKneeHandle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func drawKneeHandle(w http.ResponseWriter, r *http.Request) {
	knee := Knee {
		id: "inital knee",
		hip: Hip {
			height: utils.GetUrlIntParam(r, "hip_height", 100),
			power:  utils.GetUrlIntParam(r, "hip_power", 100),
		},
		shin: Shin {
			height: utils.GetUrlIntParam(r, "shin_height", 100),
			power:  utils.GetUrlIntParam(r, "shin_power", 100),
		},
		joint: Joint {
			degree: utils.GetUrlIntParam(r, "joint_degree", 100),
		},
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500,500)

	//hip
	s.Line(250,250, int(250-knee.hip.height), int(250-knee.hip.height), "stroke:black;stroke-width:6;")
	//shin
	s.Line(250,250, int(250+knee.shin.height), int(250+knee.shin.height*2), "stroke:red;stroke-width:6;")

	//joint
	s.Circle(250,250, 10, "fill:black")

	s.End()
}
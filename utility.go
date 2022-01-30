package main

func GetDristTypeFromFn(fname string) DristType {
	ext := fname[len(fname)-4:]
	if ext == ".jpg" {
		return Photo
	} else if ext == ".gif" {
		return Animation
	} else if ext == ".mp4" {
		return Video
	} else {
		return None
	}
}

func modFilenameForList(fname string) string {
	drtype := GetDristTypeFromFn(fname)
	drname := fname[:len(fname)-4]
	switch drtype {
	case Photo:
		return "PIC " + drname
	case Animation:
		return "GIF " + drname
	case Video:
		return "VID " + drname
	default:
		return "WTF " + drname
	}
}

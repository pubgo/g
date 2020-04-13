package httputil

import "strings"

func ExtToContentType(extName string) string {
	switch strings.ToLower(extName) {
	case "ez":
		return "application/andrew-inset"

	case "hqx":
		return "application/mac-binhex40"

	case "cpt":
		return "application/mac-compactpro"

	case "doc":
		return "application/msword"

	case "bin", "dms", "lha", "lzh", "exe", "class", "so", "dll":
		return "application/octet-stream"

	case "oda":
		return "application/oda"

	case "pdf":
		return "application/pdf"

	case "ai":
		return "application/postscript"

	case "eps":
		return "application/postscript"

	case "ps":
		return "application/postscript"

	case "smi", "smil":
		return "application/smil"

	case "mif":
		return "application/vnd.mif"

	case "xls":
		return "application/vnd.ms-excel"

	case "ppt":
		return "application/vnd.ms-powerpoint"

	case "wbxml":
		return "application/vnd.wap.wbxml"

	case "wmlc":
		return "application/vnd.wap.wmlc"

	case "wmlsc":
		return "application/vnd.wap.wmlscriptc"

	case "bcpio":
		return "application/x-bcpio"

	case "vcd":
		return "application/x-cdlink"

	case "pgn":
		return "application/x-chess-pgn"

	case "cpio":
		return "application/x-cpio"

	case "csh":
		return "application/x-csh"

	case "dcr":
		return "application/x-director"

	case "dir":
		return "application/x-director"

	case "dxr":
		return "application/x-director"

	case "dvi":
		return "application/x-dvi"

	case "spl":
		return "application/x-futuresplash"

	case "gtar":
		return "application/x-gtar"

	case "hdf":
		return "application/x-hdf"

	case "js":
		return "application/x-javascript"

	case "skp", "skd", "skt", "skm":
		return "application/x-koan"

	case "latex":
		return "application/x-latex"

	case "nc":
		return "application/x-netcdf"

	case "cdf":
		return "application/x-netcdf"

	case "sh":
		return "application/x-sh"

	case "shar":
		return "application/x-shar"

	case "swf":
		return "application/x-shockwave-flash"

	case "sit":
		return "application/x-stuffit"

	case "sv4cpio":
		return "application/x-sv4cpio"

	case "sv4crc":
		return "application/x-sv4crc"

	case "tar":
		return "application/x-tar"

	case "tcl":
		return "application/x-tcl"

	case "tex":
		return "application/x-tex"

	case "texinfo", "texi":
		return "application/x-texinfo"

	case "t", "tr", "roff":
		return "application/x-troff"

	case "man":
		return "application/x-troff-man"

	case "me":
		return "application/x-troff-me"

	case "ms":
		return "application/x-troff-ms"

	case "ustar":
		return "application/x-ustar"

	case "src":
		return "application/x-wais-source"

	case "xhtml", "xht":
		return "application/xhtml+xml"

	case "zip":
		return "application/zip"

	case "au", "snd":
		return "audio/basic"

	case "mid", "midi", "kar":
		return "audio/midi"

	case "mpga", "mp2", "mp3":
		return "audio/mpeg"

	case "aif", "aiff", "aifc":
		return "audio/x-aiff"

	case "m3u":
		return "audio/x-mpegurl"

	case "ram", "rm", "":
		return "audio/x-pn-realaudio"

	case "rpm":
		return "audio/x-pn-realaudio-plugin"

	case "ra":
		return "audio/x-realaudio"

	case "wav":
		return "audio/x-wav"

	case "pdb":
		return "chemical/x-pdb"

	case "xyz":
		return "chemical/x-xyz"

	case "bmp":
		return "image/bmp"

	case "gif":
		return "image/gif"

	case "ief":
		return "image/ief"

	case "jpeg", "jpg", "jpe":
		return "image/jpeg"

	case "png":
		return "image/jpeg"

	case "tiff", "tif":
		return "image/tiff"

	case "djvu", "djv":
		return "image/vnd.djvu"

	case "wbmp":
		return "image/vnd.wap.wbmp"

	case "ras":
		return "image/x-cmu-raster"

	case "pnm":
		return "image/x-portable-anymap"

	case "pbm":
		return "image/x-portable-bitmap"

	case "pgm":
		return "image/x-portable-graymap"

	case "ppm":
		return "image/x-portable-pixmap"

	case "rgb":
		return "image/x-rgb"

	case "xbm", "xpm":
		return "image/x-xbitmap"

	case "xwd":
		return "image/x-xwindowdump"

	case "igs":
		return "model/iges"

	case "iges":
		return "model/iges"

	case "msh", "mesh", "silo":
		return "model/mesh"

	case "wrl", "vrml":
		return "model/vrml"

	case "css":
		return "text/css"

	case "html", "htm":
		return "text/html"

	case "asc", "txt":
		return "text/plain"

	case "rtx":
		return "text/richtext"

	case "rtf":
		return "text/rtf"

	case "sgml":
		return "text/sgml"

	case "sgm":
		return "text/sgml"

	case "tsv":
		return "text/tab-separated-values"

	case "wml":
		return "text/vnd.wap.wml"

	case "wmls":
		return "text/vnd.wap.wmlscript"

	case "etx":
		return "text/x-setext"

	case "xsl", "xml":
		return "text/xml"

	case "mpeg", "mpg", "mpe":
		return "video/mpeg"

	case "qt", "mov":
		return "video/quicktime"

	case "mxu":
		return "video/vnd.mpegurl"

	case "avi":
		return "video/x-msvideo"

	case "movie":
		return "video/x-sgi-movie"

	case "ice":
		return "x-conference/x-cooltalk"

	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"

	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	case "pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"

	default:
		return "text/plain"
	}
}

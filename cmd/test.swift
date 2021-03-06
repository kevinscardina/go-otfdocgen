
//
// generated by otfdocgen, a go-lang OTF icon font generator
//

import UIKit
    
/**
 Different Glyphs for runes in OTFIcons 

 [OTFDocGen](https://github.com/kevinscardina/swift-otficon)
 */
enum OTFIcons: String, CustomStringConvertible, CaseIterable {
	/// The OTFIcons as a UIFont
	static func font(size:CGFloat) -> UIFont? {
		return UIFont(name: "OTFIcons", size: size)
	}

	 case icon0020 = "\u{0020}"
	 case icon0030 = "\u{0030}"
	 case icon0031 = "\u{0031}"
	 case icon0032 = "\u{0032}"
	 case icon0033 = "\u{0033}"
	 case icon0034 = "\u{0034}"
	 case icon0035 = "\u{0035}"
	 case icon0036 = "\u{0036}"
	 case icon0041 = "\u{0041}"
	 case icon0042 = "\u{0042}"
	 case icon0043 = "\u{0043}"
	 case icon0044 = "\u{0044}"
	 case icon0045 = "\u{0045}"
	 case icon0046 = "\u{0046}"
	 case icon0047 = "\u{0047}"
	 case icon0048 = "\u{0048}"
	 case icon0049 = "\u{0049}"
	 case icon004A = "\u{004A}"
	 case icon004B = "\u{004B}"
	 case icon004C = "\u{004C}"
	 case icon004D = "\u{004D}"
	 case icon004E = "\u{004E}"
	 case icon0050 = "\u{0050}"
	 case icon0051 = "\u{0051}"
	 case icon0052 = "\u{0052}"
	 case icon0053 = "\u{0053}"
	 case icon0055 = "\u{0055}"
	 case icon0056 = "\u{0056}"
	 case icon0057 = "\u{0057}"
	 case icon0061 = "\u{0061}"
	 case icon0062 = "\u{0062}"
	 case icon0063 = "\u{0063}"
	 case icon0064 = "\u{0064}"
	 case icon0065 = "\u{0065}"
	 case icon0066 = "\u{0066}"
	 case icon0067 = "\u{0067}"
	 case icon0068 = "\u{0068}"
	 case icon0069 = "\u{0069}"
	 case icon006A = "\u{006A}"
	 case icon006B = "\u{006B}"
	 case icon006C = "\u{006C}"
	 case icon006D = "\u{006D}"
	 case icon006E = "\u{006E}"
	 case icon006F = "\u{006F}"
	 case icon0070 = "\u{0070}"
	 case icon007A = "\u{007A}"
	

	/// Gets a String of the icon
	var string: String? {
		return self.rawValue
	}
    
	/// Convenience
	var description: String {
		return string ?? "?"
	}
}

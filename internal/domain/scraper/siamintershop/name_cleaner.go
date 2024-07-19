package siamintershop

func GetNameCleanerPattern() []string {
	return []string{
		`\[NEW\]`, `\[แยกเล่ม\]`, `\(ไม่แต้มสี\)`, `\(รวมเซ็ท\)`, `\(เล่มเดียวจบ\)`,
		`\(จบ\)`, `\(เล่มจบ\)`, `\(comic\)`, `\(แพ็คชุด\)`, `\(นิยาย\)`, `\(ฉบับนิยาย\)`, `\(ภาคพิเศษ\)`,
		`ลดจ\.`, `\(เล่มเดียวจบ\)`, `\([0-9,]+\.-\)`, `เล่ม \d+-\d+`, `เล่ม \d+`, `- \d\d`, `\d\d-\d\d`,
		`\+15.5`, `\+ Movie`, ` เล่มเดียวจบ`,
	}
}

type NameCleaner interface {
	Clean(string) string
}

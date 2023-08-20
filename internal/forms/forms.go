package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// 創建並初始化一個 Form 結構的實例。
// 將表單數據和一個空的錯誤集合關聯起來。
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// 檢查一個或多個指定的標籤（tag）是否具有必填值。
func (f *Form) HasRequired(tagIDs ...string) {
	for _, tagID := range tagIDs {
		value := f.Get(tagID)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddError(tagID, "This field can't be blank")
		}
	}
}

// 檢查特定標籤在 HTTP 請求的表單數據中是否具有值。
func (f *Form) HasValue(tagID string, r *http.Request) bool {
	x := r.Form.Get(tagID)
	// if x == "" {
	// 	f.Errors.AddError(tagID, "Field Empty")
	// 	return false
	// }
	return x != ""
}

// 檢查特定標籤在 HTTP 請求的表單數據中的值是否至少具有指定的最小長度。
func (f *Form) MinLength(tagID string, length int, r *http.Request) bool {
	x := r.Form.Get(tagID)
	if len(x) < length {
		f.Errors.AddError(tagID, fmt.Sprintf("This field must be %d charcaters long or more", length))
		return false
	}
	return true
}

// 驗證特定標籤在表單數據中的值是否為有效的電子郵件地址。
func (f *Form) IsEmail(tagID string) {
	if !govalidator.IsEmail(f.Get(tagID)) {
		f.Errors.AddError(tagID, "Invaid Email")
	}
}

// 檢查整個表單是否有效，即是否存在錯誤。
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

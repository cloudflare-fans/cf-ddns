package bu_type

import "time"

type GetDNSRecordRespBody struct {
	Errors     []interface{} `json:"errors"`
	Messages   []interface{} `json:"messages"`
	Success    bool          `json:"success"`
	ResultInfo struct {
		Count      int `json:"count"`
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Result []struct {
		Comment  string `json:"comment"`
		Name     string `json:"name"`
		Proxied  bool   `json:"proxied"`
		Settings struct {
		} `json:"settings"`
		Tags              []interface{} `json:"tags"`
		Ttl               int           `json:"ttl"`
		Content           string        `json:"content"`
		Type              string        `json:"type"`
		CommentModifiedOn time.Time     `json:"comment_modified_on"`
		CreatedOn         time.Time     `json:"created_on"`
		Id                string        `json:"id"`
		Meta              struct {
		} `json:"meta"`
		ModifiedOn     time.Time `json:"modified_on"`
		Proxiable      bool      `json:"proxiable"`
		TagsModifiedOn time.Time `json:"tags_modified_on"`
	} `json:"result"`
}

type PutDNSRecordRespBody struct {
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Success  bool          `json:"success"`
	Result   struct {
		Comment  string `json:"comment"`
		Name     string `json:"name"`
		Proxied  bool   `json:"proxied"`
		Settings struct {
		} `json:"settings"`
		Tags              []interface{} `json:"tags"`
		Ttl               int           `json:"ttl"`
		Content           string        `json:"content"`
		Type              string        `json:"type"`
		CommentModifiedOn time.Time     `json:"comment_modified_on"`
		CreatedOn         time.Time     `json:"created_on"`
		Id                string        `json:"id"`
		Meta              struct {
		} `json:"meta"`
		ModifiedOn     time.Time `json:"modified_on"`
		Proxiable      bool      `json:"proxiable"`
		TagsModifiedOn time.Time `json:"tags_modified_on"`
	} `json:"result"`
}

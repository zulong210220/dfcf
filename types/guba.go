package types

import "time"

type DfcfGubaBody struct {
	QMarket           int           `json:"q_market"`
	QCode             string        `json:"q_code"`
	Re                []*GubaRe     `json:"re"`
	Ret               []interface{} `json:"ret"`
	RetAd             []interface{} `json:"ret_ad"`
	RetAdType         []interface{} `json:"ret_ad_type"`
	Count             int           `json:"count"`
	BarName           string        `json:"bar_name"`
	BarCode           string        `json:"bar_code"`
	BarQuotCode       string        `json:"bar_quot_code"`
	Apis              int           `json:"apis"`
	BarRank           int           `json:"bar_rank"`
	StockbarFansCount int           `json:"stockbar_fans_count"`
	StockbarDesc      string        `json:"stockbar_desc"`
	HotArticles       []interface{} `json:"hot_articles"`
	IntelligentReply  int           `json:"intelligent_reply"`
	ClassicReply      int           `json:"classic_reply"`
	Cachetag          interface{}   `json:"cachetag"`
	Rc                int           `json:"rc"`
	Me                string        `json:"me"`
	Time              time.Time     `json:"time"`
}

type EsGubaReply struct {
	QCode string `json:"q_code"`
	*GubaRe
	//Re    *GubaRe `json:"re"`
}

type GubaRe struct {
	PostContent             string        `json:"post_content"`
	PostAbstract            string        `json:"post_abstract"`
	PostPublishTime         string        `json:"post_publish_time"`
	PostDisplayTime         string        `json:"post_display_time"`
	PostIP                  string        `json:"post_ip"`
	PostState               int           `json:"post_state"`
	PostCheckState          int           `json:"post_checkState"`
	PostForwardCount        int           `json:"post_forward_count"`
	PostCommentAuthority    int           `json:"post_comment_authority"`
	PostLikeCount           int           `json:"post_like_count"`
	PostIsLike              bool          `json:"post_is_like"`
	PostIsCollected         bool          `json:"post_is_collected"`
	PostType                int           `json:"post_type"`
	PostSourceID            string        `json:"post_source_id"`
	PostTopStatus           int           `json:"post_top_status"`
	PostStatus              int           `json:"post_status"`
	PostFrom                string        `json:"post_from"`
	PostFromNum             int           `json:"post_from_num"`
	PostPdfURL              string        `json:"post_pdf_url"`
	PostHasPic              bool          `json:"post_has_pic"`
	HasPicNotIncludeContent bool          `json:"has_pic_not_include_content"`
	PostPicURL              []interface{} `json:"post_pic_url"`
	SourcePostID            int           `json:"source_post_id"`
	SourcePostState         int           `json:"source_post_state"`
	SourcePostUserID        string        `json:"source_post_user_id"`
	SourcePostUserNickname  string        `json:"source_post_user_nickname"`
	SourcePostUserType      int           `json:"source_post_user_type"`
	SourcePostUserIsMajia   bool          `json:"source_post_user_is_majia"`
	SourcePostPicURL        []interface{} `json:"source_post_pic_url"`
	SourcePostTitle         string        `json:"source_post_title"`
	SourcePostContent       string        `json:"source_post_content"`
	SourcePostIP            string        `json:"source_post_ip"`
	SourcePostType          int           `json:"source_post_type"`
	SourcePostGuba          struct {
		StockbarType         int         `json:"stockbar_type"`
		StockbarCode         string      `json:"stockbar_code"`
		StockbarInnerCode    interface{} `json:"stockbar_inner_code"`
		StockbarName         string      `json:"stockbar_name"`
		StockbarMarket       string      `json:"stockbar_market"`
		StockbarQuote        int         `json:"stockbar_quote"`
		StockbarExchange     int         `json:"stockbar_exchange"`
		StockbarExternalCode string      `json:"stockbar_external_code"`
	} `json:"source_post_guba"`
	PostVideoURL        interface{} `json:"post_video_url"`
	SourcePostVideoURL  interface{} `json:"source_post_video_url"`
	SourcePostFrom      string      `json:"source_post_from"`
	SourcePostLikeCount int         `json:"source_post_like_count"`
	CodeName            string      `json:"code_name"`
	ProductType         string      `json:"product_type"`
	VUserCode           string      `json:"v_user_code"`
	SourceClickCount    interface{} `json:"source_click_count"`
	SourceCommentCount  string      `json:"source_comment_count"`
	SourceForwardCount  string      `json:"source_forward_count"`
	SourcePublishTime   string      `json:"source_publish_time"`
	SourceUserIsMajia   string      `json:"source_user_is_majia"`
	AskChairmanState    interface{} `json:"ask_chairman_state"`
	SelectedPostCode    string      `json:"selected_post_code"`
	SelectedPostName    string      `json:"selected_post_name"`
	SelectedRelateGuba  interface{} `json:"selected_relate_guba"`
	AskQuestion         string      `json:"ask_question"`
	AskAnswer           string      `json:"ask_answer"`
	Qa                  interface{} `json:"qa"`
	Extend              struct {
		DouguInfo struct {
			Relateguba struct {
				StockbarType         int    `json:"stockbar_type"`
				StockbarCode         string `json:"stockbar_code"`
				StockbarInnerCode    string `json:"stockbar_inner_code"`
				StockbarName         string `json:"stockbar_name"`
				StockbarMarket       string `json:"stockbar_market"`
				StockbarQuote        int    `json:"stockbar_quote"`
				StockbarExchange     int    `json:"stockbar_exchange"`
				StockbarExternalCode string `json:"stockbar_external_code"`
			} `json:"relateguba"`
			ClosePositionStatus int `json:"close_position_status"`
			ClosePosition       struct {
				Price   float64     `json:"price"`
				Time    interface{} `json:"time"`
				Efftime string      `json:"efftime"`
			} `json:"close_position"`
			PostTagStatus int `json:"post_tag_status"`
			PostStart     struct {
				Price float64 `json:"price"`
				Time  string  `json:"time"`
			} `json:"post_start"`
			BullishBearishTag int    `json:"bullish_bearish_tag"`
			CoverImage        string `json:"cover_image"`
			InitRankstate     int    `json:"init_rankstate"`
			Relatedincode     string `json:"relatedincode"`
		} `json:"dougu_Info"`
	} `json:"extend"`
	SourceExtend       interface{} `json:"source_extend"`
	SourcePostSourceID string      `json:"source_post_source_id"`
	PostPicURL2        interface{} `json:"post_pic_url2"`
	SourcePostPicURL2  interface{} `json:"source_post_pic_url2"`
	RelateTopic        struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		H5URL string `json:"h5_url"`
		Guide string `json:"guide"`
	} `json:"relate_topic"`
	ZwpageFlag             int              `json:"zwpage_flag"`
	SourcePostCommentCount int              `json:"source_post_comment_count"`
	PostAtuser             []interface{}    `json:"post_atuser"`
	ReplyList              []*GubaReplyList `json:"reply_list"`
	ContentType            int              `json:"content_type"`
	PostDiscussion         struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"post_discussion"`
	MediaType   int `json:"media_type"`
	RepostState int `json:"repost_state"`
	PostID      int `json:"post_id"`
	PostUser    struct {
		UserID           string        `json:"user_id"`
		UserNickname     string        `json:"user_nickname"`
		UserName         string        `json:"user_name"`
		UserV            int           `json:"user_v"`
		UserType         int           `json:"user_type"`
		UserIsMajia      bool          `json:"user_is_majia"`
		UserLevel        int           `json:"user_level"`
		UserFirstEnName  string        `json:"user_first_en_name"`
		UserAge          string        `json:"user_age"`
		UserInfluLevel   int           `json:"user_influ_level"`
		UserBlackType    int           `json:"user_black_type"`
		UserThirdIntro   interface{}   `json:"user_third_intro"`
		UserBizflag      string        `json:"user_bizflag"`
		UserBizsubflag   string        `json:"user_bizsubflag"`
		UserMedalDetails []interface{} `json:"user_medal_details"`
		UserExtendinfos  struct {
			UserAccreditinfos interface{} `json:"user_accreditinfos"`
			Deactive          string      `json:"deactive"`
		} `json:"user_extendinfos"`
	} `json:"post_user"`
	PostGuba struct {
		StockbarType         int    `json:"stockbar_type"`
		StockbarCode         string `json:"stockbar_code"`
		StockbarInnerCode    string `json:"stockbar_inner_code"`
		StockbarName         string `json:"stockbar_name"`
		StockbarMarket       string `json:"stockbar_market"`
		StockbarQuote        int    `json:"stockbar_quote"`
		StockbarExchange     int    `json:"stockbar_exchange"`
		StockbarExternalCode string `json:"stockbar_external_code"`
	} `json:"post_guba"`
	PostTitle        string      `json:"post_title"`
	PostLastTime     string      `json:"post_last_time"`
	PostClickCount   int         `json:"post_click_count"`
	PostCommentCount int         `json:"post_comment_count"`
	PostAddress      interface{} `json:"post_address"`
}

type GubaReplyList struct {
	ReplyID   int64 `json:"reply_id"`
	ReplyUser struct {
		UserID           string        `json:"user_id"`
		UserNickname     string        `json:"user_nickname"`
		UserName         string        `json:"user_name"`
		UserV            int           `json:"user_v"`
		UserType         int           `json:"user_type"`
		UserIsMajia      bool          `json:"user_is_majia"`
		UserLevel        int           `json:"user_level"`
		UserFirstEnName  string        `json:"user_first_en_name"`
		UserAge          string        `json:"user_age"`
		UserInfluLevel   int           `json:"user_influ_level"`
		UserBlackType    int           `json:"user_black_type"`
		UserThirdIntro   interface{}   `json:"user_third_intro"`
		UserBizflag      string        `json:"user_bizflag"`
		UserBizsubflag   string        `json:"user_bizsubflag"`
		UserMedalDetails []interface{} `json:"user_medal_details"`
		UserExtendinfos  struct {
			UserAccreditinfos interface{} `json:"user_accreditinfos"`
			Deactive          string      `json:"deactive"`
		} `json:"user_extendinfos"`
	} `json:"reply_user"`
	ReplyAr       string             `json:"reply_ar"`
	ReplyTime     string             `json:"reply_time"`
	ReplyText     string             `json:"reply_text"`
	SourceReply   []*GubaSourceReply `json:"source_reply"`
	ReplyPicture  string             `json:"reply_picture"`
	ReplyIsAuthor bool               `json:"reply_is_author"`
}

type GubaSourceReply struct {
	SourceReplyID   int64 `json:"source_reply_id"`
	SourceReplyUser struct {
		UserID           string        `json:"user_id"`
		UserNickname     string        `json:"user_nickname"`
		UserName         string        `json:"user_name"`
		UserV            int           `json:"user_v"`
		UserType         int           `json:"user_type"`
		UserIsMajia      bool          `json:"user_is_majia"`
		UserLevel        int           `json:"user_level"`
		UserFirstEnName  string        `json:"user_first_en_name"`
		UserAge          string        `json:"user_age"`
		UserInfluLevel   int           `json:"user_influ_level"`
		UserBlackType    int           `json:"user_black_type"`
		UserThirdIntro   interface{}   `json:"user_third_intro"`
		UserBizflag      string        `json:"user_bizflag"`
		UserBizsubflag   string        `json:"user_bizsubflag"`
		UserMedalDetails []interface{} `json:"user_medal_details"`
		UserExtendinfos  struct {
			UserAccreditinfos interface{} `json:"user_accreditinfos"`
			Deactive          string      `json:"deactive"`
		} `json:"user_extendinfos"`
	} `json:"source_reply_user"`
	SourceReplyAr      string `json:"source_reply_ar"`
	SourceReplyTime    string `json:"source_reply_time"`
	SourceReplyText    string `json:"source_reply_text"`
	SourceReplyPicture string `json:"source_reply_picture"`
}

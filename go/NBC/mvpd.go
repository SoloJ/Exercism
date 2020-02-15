package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var platforms = map[string]string{
	"iOS": "",
	"web": "",
}

type data struct {
	GlobalSettings brandSettings `json:"globalSettings"`
	MvpdWhitelist  []whitelist   `json:"mvpdWhitelist"`
}
type brandSettings struct {
	Brand                  string `json:"brand"`
	File_path              string `json:"file_path"`
	AdobePassEndPoint      string `json:"adobePassEndPoint"`
	AdobePassErrorMappings Errors `json:"adobePassErrorMappings"`
}
type Errors struct {
	Zero027 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"0027"`
	Zero027__The_E_token_is_corrupted struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"0027: The E-token is corrupted"`
	Zero033_The_customer_is_not_authorized_for_the_content_requested struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"0033:The customer is not authorized for the content requested"`
	Zero070 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"0070"`
	Zero070_We_updated_our_app__Just_sign_out_and_sign_back_into_DIRECTV_NOW_this_one_time__and_you_ll_be_all_set struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"0070-We updated our app. Just sign out and sign back into DIRECTV NOW this one time, and you'll be all set!"`
	Zero070__The_E_token_is_corrupted struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"0070: The E-token is corrupted"`
	Access_to_this_program_is_not_included_in_your_Spectrum_TV_subscription__To_learn_more_about_adding_this_program__visit_https___www_spectrum_net_or_call__855__757_7328 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Access to this program is not included in your Spectrum TV subscription. To learn more about adding this program, visit https://www.spectrum.net or call (855) 757-7328."`
	Authorization_failed__If_you_are_receiving_this_in_error__please_log_out_from_your_provider_and_try_again struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Authorization failed. If you are receiving this in error, please log out from your provider and try again."`
	Cannot_complete_action struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Cannot complete action"`
	Default_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Default Error"`
	Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Error"`
	Generic_Authentication_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Generic Authentication Error"`
	Generic_Authorization_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Generic Authorization Error"`
	Internal_Authentication_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Internal Authentication Error"`
	Internal_Authorization_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Internal Authorization Error"`
	Internal_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Internal Error"`
	It_appears_that_your_RCN_subscription_does_not_include_Bravo___If_you_believe_you_have_received_this_message_in_error__please_call_RCN_Customer_Care_at_1_800_746_4726 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your RCN subscription does not include Bravo.  If you believe you have received this message in error, please call RCN Customer Care at 1-800-746-4726."`
	It_appears_that_your_RCN_subscription_does_not_include_E____If_you_believe_you_have_received_this_message_in_error__please_call_RCN_Customer_Care_at_1_800_746_4726 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your RCN subscription does not include E!.  If you believe you have received this message in error, please call RCN Customer Care at 1-800-746-4726."`
	It_appears_that_your_RCN_subscription_does_not_include_NBC_Entertainment___If_you_believe_you_have_received_this_message_in_error__please_call_RCN_Customer_Care_at_1_800_746_4726 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your RCN subscription does not include NBC Entertainment.  If you believe you have received this message in error, please call RCN Customer Care at 1-800-746-4726."`
	It_appears_that_your_RCN_subscription_does_not_include_Oxygen___If_you_believe_you_have_received_this_message_in_error__please_call_RCN_Customer_Care_at_1_800_746_4726 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your RCN subscription does not include Oxygen.  If you believe you have received this message in error, please call RCN Customer Care at 1-800-746-4726."`
	It_appears_that_your_RCN_subscription_does_not_include_Syfy___If_you_believe_you_have_received_this_message_in_error__please_call_RCN_Customer_Care_at_1_800_746_4726 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your RCN subscription does not include Syfy.  If you believe you have received this message in error, please call RCN Customer Care at 1-800-746-4726."`
	It_appears_that_your_RCN_subscription_does_not_include_USA___If_you_believe_you_have_received_this_message_in_error__please_call_RCN_Customer_Care_at_1_800_746_4726 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your RCN subscription does not include USA.  If you believe you have received this message in error, please call RCN Customer Care at 1-800-746-4726."`
	It_appears_that_your_Suddenlink_subscription_does_not_include_Bravo__To_sign_up_for_Bravo__please_call_877_694_9474 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your Suddenlink subscription does not include Bravo. To sign up for Bravo, please call 877-694-9474."`
	It_appears_that_your_Suddenlink_subscription_does_not_include_E___To_sign_up_for_E___please_call_877_694_9474 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your Suddenlink subscription does not include E!. To sign up for E!, please call 877-694-9474."`
	It_appears_that_your_Suddenlink_subscription_does_not_include_Oxygen__To_sign_up_for_Oxygen__please_call_877_694_9474 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your Suddenlink subscription does not include Oxygen. To sign up for Oxygen, please call 877-694-9474."`
	It_appears_that_your_Suddenlink_subscription_does_not_include_SyFy__To_sign_up_for_SyFy__please_call_877_694_9474 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your Suddenlink subscription does not include SyFy. To sign up for SyFy, please call 877-694-9474."`
	It_appears_that_your_Suddenlink_subscription_does_not_include_USA__To_sign_up_for_USA__please_call_877_694_9474 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"It appears that your Suddenlink subscription does not include USA. To sign up for USA, please call 877-694-9474."`
	Multiple_Authentication_Requests_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Multiple Authentication Requests Error"`
	New_FP_version_needed_for_DRM_module struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"New FP version needed for DRM module"`
	Provider_not_Available_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Provider not Available Error"`
	Provider_not_Selected_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Provider not Selected Error"`
	The_E_token_is_corrupted struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"The E-token is corrupted"`
	The_Parental_Control_service_is_experiencing_technical_difficulties__Please_try_again_later struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"The Parental Control service is experiencing technical difficulties. Please try again later."`
	This_network_is_not_included_in_your_current_television_subscription__Please_contact_your_TV_provider_to_subscribe__http___my_xfinity_com_upgrade struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"This network is not included in your current television subscription. Please contact your TV provider to subscribe. http://my.xfinity.com/upgrade"`
	To_access_Bravo_content_you_must_subscribe_to_Bravo__To_subscribe_to_Bravo_or_if_you_subscribe_to_Bravo_and_received_this_message_in_error__please_call_Baldwin_Lightstream_Customer_Service_at_715_684_3346 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access Bravo content you must subscribe to Bravo. To subscribe to Bravo or if you subscribe to Bravo and received this message in error, please call Baldwin Lightstream Customer Service at 715-684-3346."`
	To_access_Bravo_content_you_must_subscribe_to_Bravo__To_subscribe_to_Bravo_or_if_you_subscribe_to_Bravo_and_received_this_message_in_error__please_call_Liberty_Cablevision_of_PR_Customer_Service_at_787_355_3535 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access Bravo content you must subscribe to Bravo. To subscribe to Bravo or if you subscribe to Bravo and received this message in error, please call Liberty Cablevision of PR Customer Service at 787-355-3535."`
	To_access_E__content_you_must_subscribe_to_E___To_subscribe_to_E__or_if_you_subscribe_to_E__and_received_this_message_in_error__please_call_Liberty_Cablevision_of_PR_Customer_Service_at_787_355_3535 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access E! content you must subscribe to E!. To subscribe to E! or if you subscribe to E! and received this message in error, please call Liberty Cablevision of PR Customer Service at 787-355-3535."`
	To_access_Oxygen_content_you_must_subscribe_to_Oxygen__To_subscribe_to_Oxygen_or_if_you_subscribe_to_Oxygen_and_received_this_message_in_error__please_call_CASSCOMM_Customer_Service_at_800_252_1799 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access Oxygen content you must subscribe to Oxygen. To subscribe to Oxygen or if you subscribe to Oxygen and received this message in error, please call CASSCOMM Customer Service at 800-252-1799."`
	To_access_Oxygen_content_you_must_subscribe_to_Oxygen__To_subscribe_to_Oxygen_or_if_you_subscribe_to_Oxygen_and_received_this_message_in_error__please_call_Liberty_Cablevision_of_PR_Customer_Service_at_787_355_3535 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access Oxygen content you must subscribe to Oxygen. To subscribe to Oxygen or if you subscribe to Oxygen and received this message in error, please call Liberty Cablevision of PR Customer Service at 787-355-3535."`
	To_access_Oxygen_content_you_must_subscribe_to_Oxygen__To_subscribe_to_Oxygen_or_if_you_subscribe_to_Oxygen_and_received_this_message_in_error__please_call_MetroNet_Customer_Service_at_877_407_3224 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access Oxygen content you must subscribe to Oxygen. To subscribe to Oxygen or if you subscribe to Oxygen and received this message in error, please call MetroNet Customer Service at 877-407-3224."`
	To_access_Syfy_content_you_must_subscribe_to_Syfy__To_subscribe_to_Syfy_or_if_you_subscribe_to_Syfy_and_received_this_message_in_error__please_call_Liberty_Cablevision_of_PR_Customer_Service_at_787_355_3535 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access Syfy content you must subscribe to Syfy. To subscribe to Syfy or if you subscribe to Syfy and received this message in error, please call Liberty Cablevision of PR Customer Service at 787-355-3535."`
	To_access_USA_content_you_must_subscribe_to_USA__To_subscribe_to_USA_or_if_you_subscribe_to_USA_and_received_this_message_in_error__please_call_Liberty_Cablevision_of_PR_Customer_Service_at_787_355_3535 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access USA content you must subscribe to USA. To subscribe to USA or if you subscribe to USA and received this message in error, please call Liberty Cablevision of PR Customer Service at 787-355-3535."`
	To_access_nbcentertainment_content_you_must_subscribe_to_nbcentertainment__To_subscribe_to_nbcentertainment_or_if_you_subscribe_to_nbcentertainment_and_received_this_message_in_error__please_call_Liberty_Cablevision_of_PR_Customer_Service_at_787_355_3535 struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"To access nbcentertainment content you must subscribe to nbcentertainment. To subscribe to nbcentertainment or if you subscribe to nbcentertainment and received this message in error, please call Liberty Cablevision of PR Customer Service at 787-355-3535."`
	User_Not_Authenticated_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"User Not Authenticated Error"`
	User_Not_Authorized_Error struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"User Not Authorized Error"`
	User_not_Authorized_Error__Your_Parental_Control_settings_do_not_allow_you_to_view_this_content struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"User not Authorized Error, Your Parental Control settings do not allow you to view this content."`
	We_are_sorry__but_it_seems_we_do_not_have_access_to_this_content_on_this_particular_application__We_are_working_hard_to_get_you_access__Please_check_back_soon struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"We are sorry, but it seems we do not have access to this content on this particular application. We are working hard to get you access. Please check back soon."`
	We_show_that_you_are_not_authorized_for_this_service__Please_visit_Optimum_com_or_contact_866_200_7192_to_see_how_you_can_get_access struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"We show that you are not authorized for this service. Please visit Optimum.com or contact 866-200-7192 to see how you can get access."`
	We_updated_our_app__Just_sign_out_and_sign_back_into_DIRECTV_NOW_this_one_time__and_you_ll_be_all_set struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"We updated our app. Just sign out and sign back into DIRECTV NOW this one time, and you'll be all set!"`
	We_re_sorry__something_didn_t_work_quite_right__Please_try_again struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"We're sorry, something didn't work quite right. Please try again."`
	You_are_not_currently_authorized_to_view_this_channel__For_more_information__visit_suddenlink_com_onyourside struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You are not currently authorized to view this channel. For more information, visit suddenlink.com/onyourside."`
	You_do_not_have_a_subscription_to_view_the_requested_content__Please_upgrade_your_programming_then_log_back_in_to_view_content struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You do not have a subscription to view the requested content. Please upgrade your programming then log back in to view content."`
	You_do_not_have_a_subscription_to_view_the_requested_content__To_upgrade_your_programming_go_to_mydish_com_programming_and_then_log_back_in_to_view_content struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You do not have a subscription to view the requested content. To upgrade your programming go to mydish.com/programming and then log back in to view content."`
	You_do_not_have_full_subscriptions_to_view_the_requested_content__Please_upgrade_your_programming_then_log_back_in_to_view_content struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You do not have full subscriptions to view the requested content. Please upgrade your programming then log back in to view content."`
	You_must_subscribe_to_the_package_that_includes_Bravo_as_part_of_your_PlayStation_Vue_TV_service_in_order_to_access_Bravo struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You must subscribe to the package that includes Bravo as part of your PlayStation Vue TV service in order to access Bravo."`
	You_must_subscribe_to_the_package_that_includes_NBC_as_part_of_your_PlayStation_Vue_TV_service_in_order_to_access_NBC struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You must subscribe to the package that includes NBC as part of your PlayStation Vue TV service in order to access NBC."`
	You_must_subscribe_to_the_package_that_includes_Oxygen_as_part_of_your_PlayStation_Vue_TV_service_in_order_to_access_Oxygen struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You must subscribe to the package that includes Oxygen as part of your PlayStation Vue TV service in order to access Oxygen."`
	You_must_subscribe_to_the_package_that_includes_SYFY_as_part_of_your_PlayStation_Vue_TV_service_in_order_to_access_SYFY struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You must subscribe to the package that includes SYFY as part of your PlayStation Vue TV service in order to access SYFY."`
	You_must_subscribe_to_the_package_that_includes_Sprout_as_part_of_your_PlayStation_Vue_TV_service_in_order_to_access_Sprout struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You must subscribe to the package that includes Sprout as part of your PlayStation Vue TV service in order to access Sprout."`
	You_must_subscribe_to_the_package_that_includes_USA_as_part_of_your_PlayStation_Vue_TV_service_in_order_to_access_USA struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"You must subscribe to the package that includes USA as part of your PlayStation Vue TV service in order to access USA."`
	Your_Parental_Control_settings_do_not_allow_you_to_view_this_content__https___customer_xfinity_com_Secure_OnlineParentalControls_aspx struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Your Parental Control settings do not allow you to view this content. https://customer.xfinity.com/Secure/OnlineParentalControls.aspx"`
	Your_TV_subscription_does_not_include_this_channel__Please_contact_your_TV_Provider_about_upgrading_your_service_to_get_access struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"Your TV subscription does not include this channel. Please contact your TV Provider about upgrading your service to get access."`
	HTTPS___customer_comcast_com_Secure_OnlineParentalControls_aspx struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"https://customer.comcast.com/Secure/OnlineParentalControls.aspx"`
	HTTPS___customer_xfinity_com_Secure_OnlineParentalControls_aspx_Your_Parental_Control_settings_do_not_allow_you_to_view_this_network struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"https://customer.xfinity.com/Secure/OnlineParentalControls.aspx Your Parental Control settings do not allow you to view this network."`
	NoAuthz struct {
		Body         string `json:"body"`
		Title        string `json:"title"`
		UseAdobeDesc int64  `json:"useAdobeDesc"`
	} `json:"noAuthz"`
}
type ErrorInfo struct {
	UseAdobeDesc int    `json:"useAdobeDesc"`
	Title        string `json:"title"`
	B1ody        string `json:"body"`
}
type whitelist struct {
	Mvpd                   string `json:"mvpd"`
	Mvpd_url               string `json:"mvpd_url"`
	Mvpd_display_name      string `json:"mvpd_display_name"`
	Tier                   int    `json:"tier"`
	AdvertisingKey         string `json:"advertisingKey"`
	AdobePassErrorMappings Errors `json:"adobePassErrorMappings"`
	ApppickerImage         string `json:"apppickerImage"`
	ApppickerImage_2x      string `json:"apppickerImage_2x"`
	ApploggedInImage       string `json:"apploggedInImage"`
	ApploggedInImage_2x    string `json:"apploggedInImage_2x"`
}

func main() {
	response, err := http.Get("http://mvpd-admin.nbcuni.com/mvpd/service/v2/entitlements?brand=nbcentertainment&instance=prod&platform=iOS")
	var f data
	var g interface{}
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(responseData), &f)
	err = json.Unmarshal([]byte(responseData), &g)
	fmt.Print(f, g)

}

//https://tutorialedge.net/golang/consuming-restful-api-with-go/

// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package constant

const (

	///ContentType
	//UserRelated.
	ContentTypeBegin = 100
	Text             = 101
	Picture          = 102
	Voice            = 103
	Video            = 104
	File             = 105
	AtText           = 106
	Merger           = 107
	Card             = 108
	Location         = 109
	Custom           = 110
	Revoke           = 111
	Typing           = 113
	Quote            = 114
	Transfer         = 123
	RedPacket        = 124

	AdvancedText = 117

	CustomNotTriggerConversation = 119
	CustomOnlineOnly             = 120
	ReactionMessageModifier      = 121
	ReactionMessageDeleter       = 122

	Common             = 200
	GroupMsg           = 201
	SignalMsg          = 202
	CustomNotification = 203

	// SysRelated.
	NotificationBegin = 1000

	FriendApplicationApprovedNotification = 1201 // add_friend_response
	FriendApplicationRejectedNotification = 1202 // add_friend_response
	FriendApplicationNotification         = 1203 // add_friend
	FriendAddedNotification               = 1204
	FriendDeletedNotification             = 1205 // delete_friend
	FriendRemarkSetNotification           = 1206 // set_friend_remark?
	BlackAddedNotification                = 1207 // add_black
	BlackDeletedNotification              = 1208 // remove_black
	FriendInfoUpdatedNotification         = 1209

	ConversationChangeNotification = 1300 // change conversation opt

	UserNotificationBegin        = 1301
	UserInfoUpdatedNotification  = 1303 // SetSelfInfoTip              = 204
	UserStatusChangeNotification = 1304
	UserNotificationEnd          = 1399
	OANotification               = 1400

	GroupNotificationBegin = 1500

	GroupCreatedNotification                 = 1501
	GroupInfoSetNotification                 = 1502
	JoinGroupApplicationNotification         = 1503
	MemberQuitNotification                   = 1504
	GroupApplicationAcceptedNotification     = 1505
	GroupApplicationRejectedNotification     = 1506
	GroupOwnerTransferredNotification        = 1507
	MemberKickedNotification                 = 1508
	MemberInvitedNotification                = 1509
	MemberEnterNotification                  = 1510
	GroupDismissedNotification               = 1511
	GroupMemberMutedNotification             = 1512
	GroupMemberCancelMutedNotification       = 1513
	GroupMutedNotification                   = 1514
	GroupCancelMutedNotification             = 1515
	GroupMemberInfoSetNotification           = 1516
	GroupMemberSetToAdminNotification        = 1517
	GroupMemberSetToOrdinaryUserNotification = 1518
	GroupInfoSetAnnouncementNotification     = 1519
	GroupInfoSetNameNotification             = 1520
	GroupNotificationEnd                     = 1599

	// Signaling
	SignalingNotificationBegin = 1600
	SignalingNotification      = 1601
	SignalingNotificationEnd   = 1649

	VoiceCall = 1
	VideoCall = 2

	SignalingInviation            = 1
	SignalingAccept               = 2
	SignalingReject               = 3
	SignalingJoin                 = 4
	SignalingCancel               = 5
	SignalingHungUp               = 6
	SignalingClose                = 7
	SignalingMicphoneStatusChange = 8
	SignalingSpeakStatusChange    = 9

	// voiceCall
	SignalingInvitedNotification               = 1602
	SignalingAcceptedNotification              = 1603
	SignalingRejectedNotification              = 1604
	SignalingJoinedNotification                = 1605
	SignalingCanceledNotification              = 1606
	SignalingHungUpNotification                = 1607
	SignalingMicphoneStatusChangedNotification = 1608
	SignalingClosedNotification                = 1609
	SignalingSpeakStatusChangedNotification    = 1610
	SignalingGroupInvitedNotification          = 1611
	SignalingSingleChatRejectedNotification    = 1612
	SignalingGroupJoinedNotification           = 1613
	SignalingSingleChatCanceledNotification    = 1614
	SignalingSingleChatClosedNotification      = 1615

	VoiceCallRoomEnabled  = 1
	VoiceCallRoomDisabled = 2

	RedPacketClaimedNotification       = 2401
	RedPacketExpiredNotification       = 2402
	RedPacketClaimedByUserNotification = 2403

	SuperGroupNotificationBegin  = 165
	SuperGroupUpdateNotification = 1651
	MsgDeleteNotification        = 1652
	SuperGroupNotificationEnd    = 1699

	ConversationPrivateChatNotification = 1701
	ConversationUnreadNotification      = 1702

	MsgRevokeNotification = 2101

	BusinessNotificationBegin = 2000
	BusinessNotification      = 2001

	TransferSuccessNotification  = 2002
	TransferFailedNotification   = 2003
	RechargeSuccessNotification  = 2004
	RechargeFailedNotification   = 2005
	WithdrawSuccessNotification  = 2006
	WithdrawFailedNotification   = 2007
	RedPacketTimeoutNotification = 2008
	TransferReceivedNotification = 2009

	WalletBindChangedNotification = 2010
	PayNotificationEnd            = 2020

	BusinessNotificationEnd = 2099

	// server
	ServerNotificationBegin = 1800

	ServerCreatedNotification                 = 1801
	ServerInfoSetNotification                 = 1802
	JoinServerApplicationNotification         = 1803
	ServerMemberQuitNotification              = 1804
	ServerApplicationAcceptedNotification     = 1805
	ServerApplicationRejectedNotification     = 1806
	ServerOwnerTransferredNotification        = 1807
	ServerMemberKickedNotification            = 1808
	ServerMemberInvitedNotification           = 1809
	ServerMemberEnterNotification             = 1810
	ServerDismissedNotification               = 1811
	ServerMemberMutedNotification             = 1812
	ServerMemberCancelMutedNotification       = 1813
	ServerMutedNotification                   = 1814
	ServerCancelMutedNotification             = 1818
	ServerMemberInfoSetNotification           = 1816
	ServerMemberSetToAdminNotification        = 1817
	ServerMemberSetToOrdinaryUserNotification = 1818
	ServerInfoSetAnnouncementNotification     = 1819
	ServerInfoSetNameNotification             = 1820
	ServerNotificationEnd                     = 1899

	ClearConversationNotification = 2101
	DeleteMsgsNotification        = 2102
	ModifyMessageNotification     = 2103

	HasReadReceipt = 2200

	NotificationEnd = 5000

	// status.
	MsgNormal  = 1
	MsgDeleted = 4

	// MsgFrom.
	UserMsgType = 100
	SysMsgType  = 200

	// SessionType.
	SingleChatType       = 1
	GroupChatType        = 2
	SuperGroupChatType   = 3
	NotificationChatType = 4
	ServerGroupChatType  = 5

	// token.
	NormalToken  = 0
	InValidToken = 1
	KickedToken  = 2
	ExpiredToken = 3

	// MultiTerminalLogin.
	DefalutNotKick = 0
	// Full-end login, but the same end is mutually exclusive.
	AllLoginButSameTermKick = 1
	// Only one of the endpoints can log in.
	SingleTerminalLogin = 2
	// The web side can be online at the same time, and the other side can only log in at one end.
	WebAndOther = 3
	// The PC side is mutually exclusive, and the mobile side is mutually exclusive, but the web side can be online at
	// the same time.
	PcMobileAndWeb = 4
	// The PC terminal can be online at the same time,but other terminal only one of the endpoints can login.
	PCAndOther = 5

	OnlineStatus  = "online"
	OfflineStatus = "offline"
	Registered    = "registered"
	UnRegistered  = "unregistered"

	Online  = 1
	Offline = 0

	// MsgReceiveOpt.
	ReceiveMessage          = 0
	NotReceiveMessage       = 1
	ReceiveNotNotifyMessage = 2
	ReceiveNotPushMessage   = 3

	NewMsgPushSettingAllowed = 1
	NewMsgPushSettingDenied  = 2

	// OptionsKey.
	IsHistory                  = "history"
	IsPersistent               = "persistent"
	IsOfflinePush              = "offlinePush"
	IsUnreadCount              = "unreadCount"
	IsConversationUpdate       = "conversationUpdate"
	IsSenderSync               = "senderSync"
	IsNotPrivate               = "notPrivate"
	IsSenderConversationUpdate = "senderConversationUpdate"
	IsSenderNotificationPush   = "senderNotificationPush"
	IsReactionFromCache        = "reactionFromCache"
	IsNotNotification          = "isNotNotification"
	IsSendMsg                  = "isSendMsg"

	// GroupStatus.
	GroupOk              = 0
	GroupBanChat         = 1
	GroupStatusDismissed = 2
	GroupStatusMuted     = 3

	// GroupType.
	NormalGroup  = 0
	SuperGroup   = 1
	WorkingGroup = 2
	ServerGroup  = 3

	GroupBaned          = 3
	GroupBanPrivateChat = 4

	// UserJoinGroupSource.
	JoinByAdmin = 1

	JoinByInvitation = 2
	JoinBySearch     = 3
	JoinByQRCode     = 4

	// Minio.
	MinioDurationTimes = 3600
	// Aws.
	AwsDurationTimes = 3600

	// callbackCommand.
	CallbackBeforeSendSingleMsgCommand                   = "callbackBeforeSendSingleMsgCommand"
	CallbackAfterSendSingleMsgCommand                    = "callbackAfterSendSingleMsgCommand"
	CallbackBeforeSendGroupMsgCommand                    = "callbackBeforeSendGroupMsgCommand"
	CallbackAfterSendGroupMsgCommand                     = "callbackAfterSendGroupMsgCommand"
	CallbackMsgModifyCommand                             = "callbackMsgModifyCommand"
	CallbackUserOnlineCommand                            = "callbackUserOnlineCommand"
	CallbackUserOfflineCommand                           = "callbackUserOfflineCommand"
	CallbackUserKickOffCommand                           = "callbackUserKickOffCommand"
	CallbackOfflinePushCommand                           = "callbackOfflinePushCommand"
	CallbackOnlinePushCommand                            = "callbackOnlinePushCommand"
	CallbackSuperGroupOnlinePushCommand                  = "callbackSuperGroupOnlinePushCommand"
	CallbackBeforeAddFriendCommand                       = "callbackBeforeAddFriendCommand"
	CallbackBeforeUpdateUserInfoCommand                  = "callbackBeforeUpdateUserInfoCommand"
	CallbackBeforeCreateGroupCommand                     = "callbackBeforeCreateGroupCommand"
	CallbackBeforeMemberJoinGroupCommand                 = "callbackBeforeMemberJoinGroupCommand"
	CallbackBeforeSetGroupMemberInfoCommand              = "CallbackBeforeSetGroupMemberInfoCommand"
	CallbackBeforeSetMessageReactionExtensionCommand     = "callbackBeforeSetMessageReactionExtensionCommand"
	CallbackBeforeDeleteMessageReactionExtensionsCommand = "callbackBeforeDeleteMessageReactionExtensionsCommand"
	CallbackGetMessageListReactionExtensionsCommand      = "callbackGetMessageListReactionExtensionsCommand"
	CallbackAddMessageListReactionExtensionsCommand      = "callbackAddMessageListReactionExtensionsCommand"

	// callback actionCode.
	ActionAllow     = 0
	ActionForbidden = 1
	// callback callbackHandleCode.
	CallbackHandleSuccess = 0
	CallbackHandleFailed  = 1

	// minioUpload.
	OtherType = 1
	VideoType = 2
	ImageType = 3

	// sendMsgStaus.
	MsgStatusNotExist = 0
	MsgIsSending      = 1
	MsgSendSuccessed  = 2
	MsgSendFailed     = 3

	MsgModifyServerRequestStatus = 1
	MsgModifyRedPacketStatus     = 2
)

const (
	WriteDiffusion = 0
	ReadDiffusion  = 1
)

const (
	UnreliableNotification    = 1
	ReliableNotificationNoMsg = 2
	ReliableNotificationMsg   = 3
)

const (
	AtAllString       = "AtAllTag"
	AtNormal          = 0
	AtMe              = 1
	AtAll             = 2
	AtAllAtMe         = 3
	GroupNotification = 4
)

var ContentType2PushContent = map[int64]string{
	Picture:   "[PICTURE]",
	Voice:     "[VOICE]",
	Video:     "[VIDEO]",
	File:      "[File]",
	Text:      "[TEXT]",
	AtText:    "[@TEXT]",
	GroupMsg:  "[GROUPMSG]]",
	Common:    "[NEWMSG]",
	SignalMsg: "[SIGNALINVITE]",
	Transfer:  "[Transfer]",
	RedPacket: "[RedPacket]",
}

var ContentType2PushContentI18n = map[int64]string{
	Picture:   "收到一条图片消息",
	Voice:     "收到一条语音消息",
	Video:     "收到一条视频消息",
	File:      "收到一条文件消息",
	Text:      "收到一条新消息",
	AtText:    "有人@了你",
	GroupMsg:  "收到一条群消息",
	Common:    "收到一条新消息",
	SignalMsg: "收到一条新消息",
	Transfer:  "收到一条转账消息",
	RedPacket: "收到一条红包消息",
}

const (
	FieldRecvMsgOpt    = 1
	FieldIsPinned      = 2
	FieldAttachedInfo  = 3
	FieldIsPrivateChat = 4
	FieldGroupAtType   = 5
	FieldEx            = 7
	FieldUnread        = 8
	FieldBurnDuration  = 9
	FieldHasReadSeq    = 10
)

const (
	AppOrdinaryUsers = 1
	AppAdmin         = 2

	GroupOwner         = 100
	GroupAdmin         = 60
	GroupOrdinaryUsers = 20

	GroupResponseAgree  = 1
	GroupResponseRefuse = -1

	FriendResponseNotHandle = 0
	FriendResponseAgree     = 1
	FriendResponseRefuse    = -1

	Male   = 1
	Female = 2
)

const (
	OperationID     = "operationID"
	OpUserID        = "opUserID"
	ConnID          = "connID"
	OpUserPlatform  = "platform"
	Token           = "token"
	RpcCustomHeader = "customHeader" // rpc中间件自定义ctx参数
	CheckKey        = "CheckKey"
	TriggerID       = "triggerID"
	RemoteAddr      = "remoteAddr"
)

const (
	BecomeFriendByImport = 1 // 管理员导入
	BecomeFriendByApply  = 2 // 申请添加
)

const (
	ApplyNeedVerificationInviteDirectly = 0 // 申请需要同意 邀请直接进
	AllNeedVerification                 = 1 // 所有人进群需要验证，除了群主管理员邀请进群
	Directly                            = 2 // 直接进群
)

const (
	GroupRPCRecvSize = 30
	GroupRPCSendSize = 30
)

const FriendAcceptTip = "You have successfully become friends, so start chatting"

func GroupIsBanChat(status int32) bool {
	if status != GroupStatusMuted {
		return false
	}
	return true
}

func GroupIsBanPrivateChat(status int32) bool {
	if status != GroupBanPrivateChat {
		return false
	}
	return true
}

const LogFileName = "OpenIM.log"

const LocalHost = "0.0.0.0"

// flag parse.
const (
	FlagPort                  = "port"
	FlagWsPort                = "ws_port"
	FlagTransferProgressIndex = "transferProgressIndex"
	FlagPrometheusPort        = "prometheus_port"
	FlagConf                  = "config_folder_path"
)

const OpenIMCommonConfigKey = "OpenIMServerConfig"

const CallbackCommand = "command"

const BatchNum = 100

// Subscribe to user constants
const (
	SubscriberUser = 1
	Unsubscribe    = 2
)

// server
const (
	DefaultCategoryType = 1
	SysCategoryType     = 2
	CustomCategoryType  = 3
)

const (
	ChatGroupMode = 1
	AppGroupMode  = 2
)

// join server verification
const (
	JoinServerNeedVerification = 0 // 需要审核
	JoinServerDirectly         = 1 // 直接进部落

	ServerResponseNotHandle = 0
	ServerResponseAgree     = 1
	ServerResponseRefuse    = -1
)

const (
	ServerSearchableDenied  = 0 // 不可以被搜索
	ServerSearchableAllowed = 1 // 可以被搜索
)

const (
	ServerInvitedDenied  = 0 // 不可以邀请
	ServerInvitedAllowed = 1 // 可以邀请
)

const (
	ServerCommunityPrivate = 0 // 部落社区私密
	ServerCommunityPublic  = 1 // 部落社区公开
)

const (
	ServerRoleTypeEveryOne   = 0
	ServerRoleTypeAdmin      = 1
	ServerRoleTypeSuperAdmin = 2
	ServerRoleTypeOwner      = 3
)

// club
const (
	ServerOwner         = 100
	ServerSuperAdmin    = 80
	ServerAdmin         = 60
	ServerOrdinaryUsers = 20
)

// serverStatus
const (
	ServerOk              = 0
	ServerBanChat         = 1
	ServerStatusDismissed = 2
	ServerStatusMuted     = 3
)

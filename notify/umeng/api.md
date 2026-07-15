# 友盟消息推送服务端接口使用文档

## 基本说明

本文所描述的API接口均基于HTTPS协议，使用UTF-8编码，消息体参数以及返回结果采用JSON格式。

请求方式

域名：https://msgapi.umeng.com
调用方式：POST
Content-Type: application/json
请求格式: https://msgapi.umeng.com/$url?sign=$mysign
请求样例: https://msgapi.umeng.com/api/send?sign=$mysign
签名(sign=$mysign)的计算方式参见[附录I](https://developer.umeng.com/docs/67966/detail/149296#h2--h-8)。

## 调用参数-Android
```json
{
    "appkey":"xx",    // 必填，应用唯一标识
    "timestamp":"xx",    // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
    "type":"xx",    // 必填，消息发送类型,其值可以为:
                        // unicast-单播
                        // listcast-列播，要求不超过500个device_token
                        // filecast-文件播，多个device_token可通过文件形式批量发送
                        // broadcast-广播
                        // groupcast-组播，按照filter筛选用户群,请参照filter参数
                        // customizedcast，通过alias进行推送，包括以下两种case:
                        // -alias:对单个或者多个alias进行推送
                        // -file_id:将alias存放到文件后，根据file_id来推送
    "device_tokens":"xx",    // 当type=unicast时,必填,表示指定的单个设备
                                     // 当type=listcast时,必填,要求不超过500个,以英文逗号分隔
    "alias_type":"xx",    // 当type=customizedcast时,必填
                                // alias的类型, alias_type可由开发者自定义,开发者在SDK中调用setAlias(alias, alias_type)时所设置的alias_type
    "alias":"xx",    // 当type=customizedcast时,选填(此参数和file_id二选一)
                        // 开发者填写自己的alias,要求不超过500个alias,多个alias以英文逗号间隔
                        // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
    "file_id":"xx",    // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
                          // 当type=customizedcast时，选填(此参数和alias二选一)
                          // file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应的alias_type必须和接口参数alias_type一致
                          // 使用文件播需要先调用文件上传接口获取file_id，参照"文件上传"
    "filter":{},    // 当type=groupcast时，filter或group_id至少填一个，filter代表用户筛选条件，如用户标签、渠道等，参考附录G
                     // filter的内容长度最大为3000B
    "group_id":"xx",    // 当type=groupcast时，filter或group_id至少填一个，group_id代表用户分群ID（详见Portal：运营-人群管理页面）
                     // 当filter和group_id同时存在时，只取group_id
    "payload":{    // 必填，JSON格式，具体消息内容(Android最大为2048B)
        "display_type":"xx",    // 必填，消息类型: notification(通知)、message(消息)
        "body":{    // 必填，消息体
                       // 当display_type=message时，body的内容只需填写custom字段
                       // 当display_type=notification时，body包含如下参数:
            "title":"xx",    // 必填，通知标题
            "text":"xx",    // 必填，通知文字描述
            "ticker":"xx",    //可选, 通知栏提示文字
            "big_body":"xxx",    //可选, 最多120个字符大文本
              
            // 消息重弹（推送专业版（Pro）高级能力，SDK v6.5.7及以上版本支持）:
            "re_pop":0,    // 可选，0：不重弹；1：重弹。默认值是0
              
            // 自定义通知图标:
            "icon":"xx",    // 可选，状态栏图标ID，R.drawable.[smallIcon]，
                                // 如果没有，默认使用应用图标
                                // 图片要求为24*24dp的图标，或24*24px放在drawable-mdpi下
                                // 注意四周各留1个dp的空白像素
            "img":"xx",    // 可选，通知栏大图标的URL链接。
                               // 该字段要求以http或者https开头，图片建议不大于100KB。
            "expand_image":"xx",    // 消息下方展示大图，支持自有通道消息展示
                                              // 厂商通道展示大图目前仅支持小米,要求图片为固定876*324px,仅处理在友盟推送后台上传的图片。如果上传的图片不符合小米的要求，则通过小米通道下发的消息不展示该图片，其他要求请参考小米推送文档[小米富媒体推送](https://dev.mi.com/console/doc/detail?pId=1278#_3_3 "小米富媒体推送")

            // 自定义通知声音:
            "sound":"xx",    // 可选，通知声音，R.raw.[sound]
                                  // 如果该字段为空，采用SDK默认的声音，即res/raw/下的
                                  // umeng_push_notification_default_sound声音文件。如果SDK默认声音文件不存在，则使用系统默认Notification提示音
             
            // 自定义通知样式:
            "builder_id": xx,    // 可选，默认为0，用于标识该通知采用的样式。使用该参数时
                                      // 开发者必须在App里面实现自定义通知栏样式
              
            //角标，当前设置支持友盟、华为、荣耀、vivo四个通道，小米、vivo和魅族系统默认支持自动+1，OPPO支持红点（需申请）
            //若同时存在set_badge和add_badge字段，只取set_badge的值。下述数值范围1~99仅针对华为/荣耀厂商通道；友盟通道数值不限制（如：可传-1，角标减1）
            "set_badge":5,     //可选，没有默认值。角标设置数字，范围为1~99。如果设置的值不在此区间该参数值将被忽略，需配合main_activity使用，具体说明参考main_activity
            "add_badge":1,     //可选，没有默认值。角标设置数字，范围为1~99。如果设置的值不在此区间该参数值将被忽略，需配合main_activity使用，具体说明参考main_activity
            
            // 通知到达设备后的提醒方式(注意，"true/false"为字符串):
            "play_vibrate":"true/false",    // 可选，收到通知是否震动，默认为"true"
            "play_lights":"true/false",    // 可选，收到通知是否闪灯，默认为"true"
            "play_sound":"true/false",    // 可选，收到通知是否发出声音，默认为"true"
            
            //点击"通知"的后续行为(默认为打开app):
            "after_open":"xx",    // 可选，默认为"go_app"，值可以为:
                                         // "go_app":打开应用
                                         // "go_url":跳转到URL
                                         // "go_activity":打开特定的activity
                                         // "go_custom":用户自定义内容
            "url":"xx",    // 当after_open=go_url时，必填
                             // 通知栏点击后跳转的URL，要求以http或者https开头
            "activity":"xx",    //当after_open=go_activity时，必填。
                                    // 通知栏点击后打开的Activity
            "custom":"xx"/{},    // 当display_type=message时,必填
                                       // 当display_type=notification且after_open=go_custom时，必填
                                       // 用户自定义内容，可以为字符串或者JSON格式。
        },
        
        "extra":{    // 可选，JSON格式，用户自定义key-value。
                     // 可以配合消息到达后，打开App/URL/Activity使用
            "key1":"value1",
            "key2":"value2",
            ...
        }
    },
    "policy":{    // 可选，发送策略
        "start_time":"xx",    // 可选，定时发送时，若不填写表示立即发送
                                // 定时发送时间不能小于当前时间
                                // 格式:"yyyy-MM-dd HH:mm:ss"
                                // 注意，start_time只对任务类消息生效
        "expire_time":"xx",    // 可选，消息过期时间，其值不可小于发送时间或者start_time(如果填写了的话)
                                  // 如果不填写此参数，默认为3天后过期。格式同start_time
        "max_send_num": xx,    // 可选，发送限速，每秒发送的最大条数。最小值1000
                                     //开发者发送的消息如果有请求自己服务器的资源，可以考虑此参数
        "out_biz_no":"xx" ,   // 可选，消息发送接口对任务类消息的幂等性保证
                                // 强烈建议开发者在发送任务类消息时填写这个字段，友盟服务端会根据这个字段对消息做去重避免重复发送
                                // 同一个appkey下面的多个消息会根据out_biz_no去重，不同发送任务的out_biz_no需要保证不同，否则会出现后发消息被去重过滤的情况
                                // 注意，out_biz_no只对任务类消息有效
         "uapp_arg":"xxxxx",   //可选，在uapp智能运营后台创建智能运营计划使用的计划id，参考文档 https://developer.umeng.com/docs/119267/detail/3018682
         "notification_closed_filter":true,  //可选，只对display_type=notification的消息生效，设置为true会过滤关闭通知栏消息的设备，以免占用厂商额度。
         "in_app":{
            "in_app":true       //可选，如果配置为true,在通知栏关闭时会通过应用内弹窗展示。该功能为推送专业版（Pro）高级能力，SDK v6.6.3及以上版本支持    
         }
         
    },
    "production_mode":"true/false",    // 可选，true正式模式，false测试模式。默认为true
                                                     // 广播、组播下的测试模式只会将消息发给测试设备。测试设备需要到web上添加
                                                     // 单播、文件播不受测试设备限制
    "description":"xx",    // 可选，发送消息描述，建议填写
    "category":0,    // 可选，友盟消息自分类，0：资讯营销类消息，1：服务与通讯类消息。有利于小米和vivo通道选择，建议填写。
    "callback_params": { // 自定义回执参数
        "name": "string",
        "age": 26
    },
    "channel_properties":{    // 可选，厂商通道相关的特殊配置
        "channel_activity":"xxx",  //系统弹窗，走厂商通道时必填。只有display_type=notification时有效，表示华为、小米、oppo、vivo、魅族的设备离线时走系统通道下发时打开指定页面acitivity的完整包路径。
        "xiaomi_channel_id":"",    // 小米channel_id，具体使用及限制请参考小米推送文档 https://dev.mi.com/console/doc/detail?pId=2086
        "vivo_category":"xx",      // vivo消息二级分类参数：友盟侧只进行参数透传，不做合法性校验，具体使用及限制参考[vivo消息推送分类功能说明]https://dev.vivo.com.cn/documentCenter/doc/359
        "vivo_addbadge":"true" , //vivo角标功能，需要客户端先完成系统API接入(https://dev.vivo.com.cn/documentCenter/doc/787),否则发消息时会报错，错误码为10089
        "oppo_channel_id":"xx" ,   // 可选， android8以上推送消息需要新建通道，否则消息无法触达用户。push sdk 6.0.5及以上创建了默认的通道:upush_default，消息提交厂商通道时默认添加该通道。如果要自定义通道名称或使用私信，请自行创建通道，推送消息时携带该参数具体可参考[oppo推送私信通道申请] https://open.oppomobile.com/new/developmentDoc/info?id=11227
        "oppo_category":"xx", //可选，OPPO消息类别，参考OPPO官方文档 https://open.oppomobile.com/new/developmentDoc/info?id=13189 
        "oppo_notify_level":"xx", //OPPO通知栏消息提醒等级，1-通知栏，2-通知栏+锁屏 16-通知栏+锁屏+横幅+震动+铃声，使用该参数时，oppo_category必传
        "main_activity":"xx",         // 可选，应用入口Activity类全路径,主要用于华为通道角标展示。具体使用可参考[华为角标使用说明]https://developer.umeng.com/docs/67966/detail/272597
        "huawei_channel_importance":"NORMAL",// 可选，华为&荣耀消息分类 LOW：资讯营销类消息，NORMAL：服务与通讯类消息
        "huawei_channel_category":"MARKETING", // 可选，华为自分类消息类型 [华为消息分类]https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-priority-0000001181716924
        "channel_fcm":"0",  // 可选，fcm通道开关，0不使用，1使用
        "meizu_msg_type":"0", // 可选，魅族消息自分类参数， 0：公信   1：私信，默认为公信 。可参考魅族官方文档 https://open-res.flyme.cn/fileserver/upload/file/202504/13dc49a671b643438431b386e071f59e.pdf
        "oppo_private_msg_template":{
        //oppo私信模版，请按照申请的私信模版来填充相关字段，可参考oppo私信模版相关文档 https://open.oppomobile.com/documentation/page/info?id=12391
            "private_msg_template_id":"xxxx", //oppo私信模版id
            "private_title_parameters":{"city":"北京"},      //标题模版填充参数，json格式， 若无可不填
            "private_content_parameters":{"userName":"汤姆","city":"深圳市"}  //内容模版填充参数，若无可不填                         //
        },
        "xiaomi_extra_properties":{
            //小米的extra参数，包括但不限于下述两个，可参考小米推送文档 https://dev.mi.com/xiaomihyperos/documentation/detail?pId=1559
            "template_id" : "xxxx", //发私信必填，消息模板id
            "template_param":"{}"   //消息模板参数，值为消息模板参数的JSON字符串.可参考小米推送模版接入指南 https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2314
        }
    },
    "local_properties":{    //可选，本地通知相关的特殊配置
        //请严格按照华为官方文档https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section1085395991513设置category和importance取值。
        //若category参数和importance参数均不为空，友盟只校验参数值是否合法，不校验参数的对应关系是否正确。
        //若category参数不为空，importance参数未设置，友盟会校验category参数值是否合法，同时根据category参数来智能适配对应的importance值。
        //若category参数未设置，importance参数不为空，友盟会同时忽略category参数和importance参数。
        "category":"CATEGORY_PROMO",    //可选，华为本地通知category
        "importance":"IMPORTANCE_MIN",    //可选，华为本地通知importance
        //可参考vivo关于本地通知接入的说明(https://dev.vivo.com.cn/documentCenter/doc/930)来申请channel_id和channel_name,并在发消息时增加字段
        "channel_id":"",                // 通知渠道id
        "channel_name":""               // 通知渠道名称
    } 
}
```

## 调用参数-iOS

```json
{
    "appkey":"xx",    // 必填，应用唯一标识
    "timestamp":"xx",    // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
    "type":"xx",    // 必填，消息发送类型,其值可以为: 
                        // unicast-单播
                        // listcast-列播，要求不超过500个device_token
                        // filecast-文件播，多个device_token可通过文件形式批量发送
                        // broadcast-广播
                        // groupcast-组播，按照filter筛选用户群, 请参照filter参数
                        // customizedcast，通过alias进行推送，包括以下两种case:
                        // -alias: 对单个或者多个alias进行推送
                        // -file_id: 将alias存放到文件后，根据file_id来推送
    "device_tokens":"xx",    // 当type=unicast时, 必填, 表示指定的单个设备
                                      // 当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
    "alias_type":"xx",    // 当type=customizedcast时, 必填
                                // alias的类型, alias_type可由开发者自定义, 开发者在SDK中调用setAlias(alias, alias_type)时所设置的alias_type
    "alias":"xx",    // 当type=customizedcast时, 选填(此参数和file_id二选一)
                        // 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
                        // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
    "file_id":"xx",    // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
                          // 当type=customizedcast时，选填(此参数和alias二选一)
                          // file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应的alias_type必须和接口参数alias_type一致。
                          // 使用文件播需要先调用文件上传接口获取file_id，参照"2.4文件上传接口"
    "filter":{},    // 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等，参考附录G
    "payload":{    // 必填，JSON格式，具体消息内容(iOS最大为2012B)
        "aps":{    // 必填，严格按照APNs定义来填写
            "alert":""或者{,    // 当content-available=1时(静默推送)，可选; 否则必填
                                // 可为字典类型和字符串类型
                  "title":"title",
                  "subtitle":"subtitle",
                  "body":"body"
             },
            "badge":"xx",    // 可选，取值为N（代表设置角标为N）、+N（代表角标原有基础上+N）、-N（代表角标原有基础上-N）、空字符串（代表清空角标，同N=0）。
            "sound":"xx",    // 可选         
            "content-available":1,    // 可选，代表静默推送     
            "category":"xx",    // 可选，注意: iOS8才支持该字段
            "thread-id":"xx",  // 可选，分组折叠，设置UNNotificationContent的threadIdentifier属性
            "interruption-level": "active" //可选，消息的打扰级别，iOS15起支持，四个选项"passive", "active", "time-sensitive", "critical"
        },
        "key1":"value1",    // 可选，用户自定义内容, "d","p"为友盟保留字段,key不可以是"d","p"
        "key2":"value2",
    ...
    },
    "callback_params": { // 自定义回执参数
        "name": "string",
        "age": 26
    },
    "policy":{    // 可选，发送策略
        "start_time":"xx",    // 可选，定时发送时间，若不填写表示立即发送
                                    // 定时发送时间不能小于当前时间
                                    // 格式: "yyyy-MM-dd HH:mm:ss"
                                    // 注意，start_time只对任务生效
        "expire_time":"xx",    // 可选，消息过期时间，其值不可小于发送时间或者
                                      // start_time(如果填写了的话)
                                      // 如果不填写此参数，默认为3天后过期。格式同start_time
        "out_biz_no":"xx",    // 可选，消息发送接口对任务类消息的幂等性保证
                                     // 强烈建议开发者在发送任务类消息时填写这个字段，友盟服务端会根据这个字段对消息做去重避免重复发送
                                     // 同一个appkey下面的多个消息会根据out_biz_no去重，不同发送任务的out_biz_no需要保证不同，否则会出现后发消息被去重过滤的情况
                                     // 注意，out_biz_no只对任务类消息有效
        "apns_collapse_id":"xx",    // 可选，多条带有相同apns_collapse_id的消息，iOS设备仅展示
                                            // 最新的一条，字段长度不得超过64bytes
        "uapp_arg":"xxxxx",   //可选，在uapp智能运营后台创建智能运营计划使用的计划id，只能发正式消息，参考文档 https://developer.umeng.com/docs/119267/detail/3018682
    },
    "production_mode":"true/false",    // 可选，正式/测试模式。默认为true
                                                    // 广播、组播下的测试模式只会将消息发给测试设备。测试设备需要到web上添加
                                                     // 单播、文件播不受测试设备限制
    "description":"xx"    // 可选，发送消息描述，建议填写接口     
}
```

## 调用参数-Harmony

### 普通消息
```json
{
    "appkey": "xx",    // 必填，应用唯一标识
    "timestamp": "xx",    // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
    "type": "xx",    // 必填，消息发送类型,其值可以为:
                        // unicast-单播
                        // listcast-列播，要求不超过500个device_token
                        // filecast-文件播，多个device_token可通过文件形式批量发送
                        // broadcast-广播
                        // groupcast-组播，按照filter筛选用户群,请参照filter参数
                        // customizedcast，通过alias进行推送，包括以下两种case:
                        // -alias:对单个或者多个alias进行推送
                        // -file_id:将alias存放到文件后，根据file_id来推送
    "device_tokens": "xx",    // 当type=unicast时,必填,表示指定的单个设备
                                     // 当type=listcast时,必填,要求不超过500个,以英文逗号分隔
    "alias_type": "xx",    // 当type=customizedcast时,必填
                                // alias的类型, alias_type可由开发者自定义,开发者在SDK中调用setAlias(alias, alias_type)时所设置的alias_type
    "alias": "xx",    // 当type=customizedcast时,选填(此参数和file_id二选一)
                        // 开发者填写自己的alias,要求不超过500个alias,多个alias以英文逗号间隔
                        // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
    "file_id": "xx",    // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
                          // 当type=customizedcast时，选填(此参数和alias二选一)
                          // file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应的alias_type必须和接口参数alias_type一致
                          // 使用文件播需要先调用文件上传接口获取file_id，参照"文件上传"
    "filter": {},    // 当type=groupcast时，filter或group_id至少填一个，filter代表用户筛选条件，如用户标签、渠道等，参考附录G
                     // filter的内容长度最大为3000B
    "group_id":"xx",    // 当type=groupcast时，filter或group_id至少填一个，group_id代表用户分群ID（详见Portal：运营-人群管理页面）
                     // 当filter和group_id同时存在时，只取group_id
    "payload": {    // 必填，JSON格式，具体消息内容(最大为2048B)
        "display_type":"xx",    // 必填，消息类型: notification(通知)、message(消息)
        "body": {    // 必填，消息体
                       // 当display_type=message时，body的内容只需填写custom字段
                       // 当display_type=notification时，body包含如下参数:
            "title":"xx",    // 必填，通知标题
            "text":"xx",    // 必填，通知文字描述
            "big_body":"xxx",    //可选, 最多120个字符大文本
              
            "large_icon":"xx",    // 可选，通知栏大图标的本地文件。
                                  // 该字段要求以图片建议不大于100KB，若超过则只展示消息，不展示图标。
                           
            //角标，当前设置支持友盟、鸿蒙
            //下述数值范围1~99仅针对鸿蒙厂商通道；友盟通道数值不限制
            "add_badge":1,     //可选，没有默认值。角标设置数字，范围为1~99。
                        
            //点击"通知"的后续行为(默认为打开app):
            "after_open":"xx",    // 可选，默认为"go_app"，值可以为:
                                         // "go_app":打开应用
                                         // "go_ability": 打开应用自定义页面。当after_open=go_ability时，字段uri和action至少填写一个。当存在多个Ability时，分别填写不同Ability的action和uri，优先使用action查找对应的应用内置页面。
            "uri": "xx",    // 应用内置页面ability对应的uri。参见鸿蒙如何设置uri: https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/push-dev-0000001727885258#ZH-CN_TOPIC_0000001727885258__p755671716257
            "action": "",   // 应用内置页面ability对应的action。参见鸿蒙如何设置action: https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/push-dev-0000001727885258#ZH-CN_TOPIC_0000001727885258__p729245812412
            "custom":"xx"/{},    // 当display_type=message时,必填
                                 // 用户自定义内容，可以为字符串或者JSON格式。
        },
        "extra": {    // 可选，JSON格式，用户自定义key-value。
                      // 可以配合消息到达后，打开App/URL/Activity使用
            "key1":"value1",
            "key2":"value2",
            ...
        }
    },
    "policy": {    // 可选，发送策略
        "start_time":"xx",    // 可选，定时发送时，若不填写表示立即发送
                                // 定时发送时间不能小于当前时间
                                // 格式:"yyyy-MM-dd HH:mm:ss"
                                // 注意，start_time只对任务类消息生效
        "expire_time":"xx",    // 可选，消息过期时间，其值不可小于发送时间或者start_time(如果填写了的话)
                                  // 如果不填写此参数，默认为3天后过期。格式同start_time
        "max_send_num": xx,    // 可选，发送限速，每秒发送的最大条数。最小值1000
                                     //开发者发送的消息如果有请求自己服务器的资源，可以考虑此参数
        "out_biz_no":"xx" ,   // 可选，消息发送接口对任务类消息的幂等性保证
                              // 强烈建议开发者在发送任务类消息时填写这个字段，友盟服务端会根据这个字段对消息做去重避免重复发送
                              // 同一个appkey下面的多个消息会根据out_biz_no去重，不同发送任务的out_biz_no需要保证不同，否则会出现后发消息被去重过滤的情况
                              // 注意，out_biz_no只对任务类消息有效
        
        /*
          通道策略，默认值为2
          1. 只通过友盟通道下发(在线下发，离线缓存消息，下次上线后发送)
          2. 在线时通过友盟通道下发，离线尝试通过厂商下发(厂商下发失败会通过友盟通道下发)
          3. 优先通过厂商通道下发，失败时通过友盟通道下发 (pro)
          4. 只通过厂商下发(失败时丢弃消息) (pro)
        */
        "channel_strategy": {
             "default": 2 ,
          }

    },
    "production_mode": "true/false",    // 可选，true正式模式，false测试模式。默认为true
                                        // 广播、组播下的测试模式只会将消息发给测试设备。测试设备需要到web上添加
                                        // 单播、文件播不受测试设备限制
    "description":"xx",    // 可选，发送消息描述，建议填写
    "callback_params": { // 自定义回执参数
        "name": "string",
        "age": 26
    },
    "channel_properties": {    // 可选，厂商通道相关的特殊配置
        "harmony_channel_category": "MARKETING" // 可选，鸿蒙消息分类类型 [鸿消息分类]
    },
}
```

### 鸿蒙场景化消息

```json
{
    "appkey": "xx",    // 必填，应用唯一标识
    "timestamp": "xx",    // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
    "type": "xx",    // 必填，消息发送类型,其值可以为: unicast-单播
    "device_tokens": "xx",    // 当type=unicast时,必填,表示指定的单个设备
    "payload": {    // 必填，JSON格式，具体消息内容(最大为2048B)
        "display_type":"xx",    // 必填，消息类型: scene（场景化消息）
        "channels_body": {
            "harmony": {
                "type": "live_view",
                "payload": "" // 厂商协议json序列化后的字符串
            },
        },
    },
    "policy": {    // 可选，发送策略
        "expire_time":"xx",    // 可选，消息过期时间，其值不可小于发送时间或者start_time(如果填写了的话)
                                  // 如果不填写此参数，默认为3天后过期。格式同start_time
    },
    "production_mode": "true/false",    // 可选，true正式模式，false测试模式。默认为true
}
```
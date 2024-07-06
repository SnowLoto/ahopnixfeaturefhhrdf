package I18n

var I18nDict_ja_JP map[uint16]string = map[uint16]string{
	ACME_FailedToGetCommand:             "ACME コマンドの解析に失敗しました。",
	ACME_FailedToSeek:                   "ファイルのシークに失敗したため、この ACME ファイルが無効のようです。",
	ACME_StructureErrorNotice:           "ファイルの内容が無効です。",
	ACME_UnknownCommand:                 "このファイルは未知な ACME コマンドを含めています。",
	Auth_BackendError:                   "バックエンドエラー",
	Auth_FailedToRequestEntry:           "サーバーの入り方の取得に失敗しました、サーバーのレベルレギュレーションをオフーにする上、パスワードが正しいかどうかを再度確認してお願いします。",
	Auth_HelperNotCreated:               "ヘルパーユーザーがまだ作成されていませんため、FastBuilder ユーザーセンターで、一度作成をお願いします。",
	Auth_InvalidFBVersion:               "このバーションはアウトデートしたので更新してください。",
	Auth_InvalidHelperUsername:          "ヘルパーユーザーがまだ名付かれていませんため、FastBuilder ユーザーセンターで名付けてください。",
	Auth_InvalidToken:                   "ログイントークンが無効です。",
	Auth_InvalidUser:                    "ユーザー無効です。",
	Auth_ServerNotFound:                 "指定されたサーバーは見つかりませんでした。開放状態を検討し、再度試してください。",
	Auth_UnauthorizedRentalServerNumber: "未認証のサーバー番号でした。FastBuilder ユーザーセンターで添加してください。",
	Auth_UserCombined:                   "指定されたユーザーはもうほかのユーザーに合体されましたので新アカウントでログインしてください。",
	Auth_FailedToRequestEntry_TryAgain:  "サーバーインフォの取得に失敗しました。あとでもう一度試して下さい。",
	BDump_Author:                        "作者",
	BDump_EarlyEOFRightWhenOpening:      "ファイルが読めませんでした。ファイルエンドが来るのは早すぎました。ファイルがもうこわれた可能性があります。",
	BDump_FailedToGetCmd1:               "cmd[pos:0] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd2:               "cmd[pos1] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd4:               "cmd[pos2] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd6:               "cmd[pos3] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd7_0:             "cmd[pos4] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd7_1:             "cmd[pos5] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd10:              "cmd[pos6] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd11:              "cmd[pos7] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetCmd12:              "cmd[pos8] のアーギュメントの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToGetConstructCmd:       "BDump コマンドの取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FailedToReadAuthorInfo:        "作者に関する情報の取得に失敗しました。ファイルがもう壊れた可能性があります。",
	BDump_FileNotSigned:                 "このファイルはサインされていません。",
	BDump_FileSigned:                    "このファイルはサインに保護されています。サイン元: %s",
	BDump_NotBDX_Invheader:              "このファイルは bdx ファイルではありません。(無効なヘッダー)",
	BDump_NotBDX_Invinnerheader:         "ファイルが bdx ファイルではありません。 (むこうな内部ヘッダー)",
	BDump_SignedVerifying:               "このファイルはサインに保護されています、ただいま検証しています...",
	BDump_VerificationFailedFor:         "%v のためにファイルのサインの検証に失敗しました。",
	BDump_Warn_Reserved:                 "警告: BDump/Import: 保留されたコマンドが使用されました\n",
	CommandNotFound:                     "未知なコマンドです。",
	ConnectionEstablished:               "サーバーへの接続に成功しました。",
	Crashed_No_Connection:               "長い時間をかけても接続できませんでした。",
	Crashed_OS_Windows:                  "ENTER を押して終了する。",
	Crashed_StackDump_And_Error:         "スタックダンプ (Stack dump) が上に見えます。エラーは: ",
	Crashed_Tip:                         "ヤバいぞ！FastBuilder Phoenix がクラッシューしたよ！",
	CurrentDefaultDelayMode:             "今のデフォルトディレーモード",
	CurrentTasks:                        "進行中のタスク:",
	DelayModeSet:                        "ディレーモードが設定されました",
	DelayModeSet_DelayAuto:              "ディレーを自動に %d に設定しました。",
	DelayModeSet_ThresholdAuto:          "ディレーのしきい値を自動に  %d に設定しました．",
	DelaySet:                            "ディレーが設定されました．",
	DelaySetUnavailableUnderNoneMode:    "[delay set] コマンドがディレーモードが none の時に使えません。",
	DelayThreshold_OnlyDiscrete:         "ディレーのしきい値がディレーモードが discrete の時だけ適用します。",
	DelayThreshold_Set:                  "ディレーのしきい値を %d に変更しました。",
	ERRORStr:                            "エラー",
	EnterPasswordForFBUC:                "FastBuilder ユーザーセンターのパスワードを入力してください: ",
	Enter_FBUC_Username:                 "FastBuilder ユーザーセンターのログイン名を入力してください: ",
	Enter_Rental_Server_Code:            "サーバー番号を入力してください: ",
	Enter_Rental_Server_Password:        "サーバーのパスワードを入力してください (パスワードがなければ直接 ENTER をおす。入力した内容がみえん): ",
	ErrorIgnored:                        "エラーを無視しました。",
	Error_MapY_Exceed:                   "3DMap で, MapY が必ず [20~255] の範囲にしてください (入力した方 = %v)",
	FBUC_LoginFailed:                    "FastBuilder ユーザーセンターのログイン名やパスワードに誤りがあります。",
	FBUC_Token_ErrOnCreate:              "トークンファイルの作成に失敗: ",
	FBUC_Token_ErrOnGen:                 "臨時トークンを作成できませんでした",
	FBUC_Token_ErrOnRemove:              "トークンファイルを削除できませんでした: %v",
	FBUC_Token_ErrOnSave:                "トークンの保存に失敗しました: ",
	FileCorruptedError:                  "ファイルが壊れました",
	Get_Warning:                         "",
	IgnoredStr:                          "無視した",
	InvalidFileError:                    "無効なファイルです。",
	InvalidPosition:                     "働けるポジションを取得できませんでした。 (無視でOK)",
	Lang_Config_ErrOnCreate:             "言語設定ファイルを作成できませんでした: %v",
	Lang_Config_ErrOnSave:               "言語設定の保存に失敗しました: %v",
	LanguageName:                        "日本語",
	LanguageUpdated:                     "言語設定を変更しました。",
	Logout_Done:                         "FastBuilder ユーザーセンターからログアウトしました。",
	Menu_BackButton:                     "< Back", // REMOVED FEATURES NO TRANSLATION
	Menu_Cancel:                         "Cancel", // FOR THEM
	Menu_CurrentPath:                    "Current path",
	Menu_ExcludeCommandsOption:          "Exclude Commands",
	Menu_GetEndPos:                      "getEndPos",
	Menu_GetPos:                         "getPos",
	Menu_InvalidateCommandsOption:       "Invalidate Commands",
	Menu_Quit:                           "Quit Program",
	Menu_StrictModeOption:               "Strict Mode",
	NotAnACMEFile:                       "このファイルは ACME ストラクチャーファイルではありませんでした。",
	Notice_CheckUpdate:                  "更新を検査中、少々お待ちして下さい...",
	Notice_iSH_Location_Service:         "iSH が適用しているため、稼働状態の保持にはロケーションサービスが必要です。位置情報が利用されませんで、いつでもオフーにしてもいいです。",
	Notice_OK:                           "完成\n",
	Notice_UpdateAvailable:              "新バーション (%s) の PhoenixBuilder がリリースされました。\n",
	Notice_UpdateNotice:                 "更新してください。\n",
	Notice_ZLIB_CVE:                     "今使われている zlib のバーション (%s) は古すぎていくつかの CVE バルネラビリティーがもう確認されました、それを更新する方がおすすめです。",
	Notify_NeedOp:                       "OP レベルが PhoenixBuilder の正常なかどうには必要です。",
	Notify_TurnOnCmdFeedBack:            "PhoenixBuilder の正常な稼働には gamerule sendcommandfeedback を true にするのが必要ですので、もう自動に　gamerule sendcommandfeedback を true にさせていただきましたので、もし必要ならばそのうち再度変更してください。",
	Omega_WaitingForOP:                  "Omega System が OP レベルの取得を待っています...",
	Omega_Enabled:                       "Omega System が起用された!",
	OpPrivilegeNotGrantedForOperation:   "OP レベルが PhoenixBuilder の正常なかどうには必要です、OP にしたらまた試して下さい。",
	Original_Important_Contributor:      "初期の重要な貢献者です: LNSSPsd(Ruphane), CAIMEO, CMA2401PT",
	Parsing_UnterminatedEscape:          "終わらず escape",
	Parsing_UnterminatedQuotedString:    "終わらず quoted string",
	PositionGot:                         "位置を取得しました。",
	PositionGot_End:                     "終点位置を取得しました。",
	PositionSet:                         "位置を設定しました",
	PositionSet_End:                     "終点位置を設定しました",
	QuitCorrectly:                       "プログラムが正常に終了しました。",
	Sch_FailedToResolve:                 "ファイルの読み込みに失敗しました",
	SelectLanguageOnConsole:             "コンソールで新言語を選択して下さい。",
	ServerCodeTrans:                     "サーバー",
	SimpleParser_Int_ParsingFailed:      "アーギュメント読み込み器: 整数アーギュメントの読み込みに失敗しました。",
	SimpleParser_InvEnum:                "アーギュメント読み込み器: 予期せず選択肢、予期した選択肢は: %s.",
	SimpleParser_Invalid_decider:        "アーギュメント読み込み器: 無効な選択肢でした",
	SimpleParser_Too_few_args:           "アーギュメント読み込み器: アーギュメント少なさ過ぎます",
	SkipMCPCheckChallengesLimit:         "「skip-mcpc-check-challenges」オプションを有効にしたため、呼び出した機能は使えません。",
	Special_Startup:                     "日本語を使用している\n",
	TaskCreated:                         "タスクを作成しました",
	TaskDisplayModeSet:                  "タスク状態表示モードを %s に変更しました．",
	TaskFailedToParseCommand:            "コマンド %v の読み込みに失敗しました。",
	TaskNotFoundMessage:                 "提供されたタスク番号で進行しているタスクを見つかりませんでした。",
	TaskPausedNotice:                    "[タスク %d] - 一時停止",
	TaskResumedNotice:                   "[タスク %d] - 回復",
	TaskStateLine:                       "ID %d - コマンドライン:\"%s\", 状態: %s, ディレー: %d, ディレーモード: %s, ディレーしきい値: %d",
	TaskStoppedNotice:                   "[タスク %d] - 停止",
	TaskTTeIuKoto:                       "タスク",
	TaskTotalCount:                      "総計: %d",
	TaskTypeCalculating:                 "計算中",
	TaskTypeDied:                        "停止",
	TaskTypePaused:                      "一時停止",
	TaskTypeRunning:                     "進行中",
	TaskTypeSpecialTaskBreaking:         "スペシャルタスク:停止中",
	TaskTypeSwitchedTo:                  "タスク作成タイプが %s に変更しました",
	TaskTypeUnknown:                     "未知",
	Task_D_NothingGenerated:             "[タスク %d] なにも建造されませんでした",
	Task_DelaySet:                       "[タスク %d] - ディレーが %d に変更しました",
	Task_ResumeBuildFrom:                "ブロック数 %d から建造過程を回復します",
	Task_SetDelay_Unavailable:           "[setdelay] がディレーモードが none の時に使用できませんです。",
	Task_Summary_1:                      "[タスク %d] %v 影響されたブロック",
	Task_Summary_2:                      "[タスク %d] %v 秒をかけた",
	Task_Summary_3:                      "[タスク %d] 平均速度: %v ブロック/秒",
	UnsupportedACMEVersion:              "このバーションの ACME ストラクチャーファイルは支援されていません。バーション 1.2 だけが支援されています。",
	Warning_ACME_Deprecated:             "WARNING - `acme' は非推奨となり、削除されました。代わりにBDX形式に移行してください。\n詳細については、https://github.com/LNSSPsd/PhoenixBuilder/issues/313 を参照してください。",
	Warning_Schem_Deprecated:            "警告 - `schem' は非推奨となり、削除されました。代わりにBDX形式に移行してください。\n詳細については、https://github.com/LNSSPsd/PhoenixBuilder/issues/313 を参照してください。",
	Warning_UserHomeDir:                 "警告 - ユーザーのホームフォルダーの取得に失敗しました。homedir=\".\";をさせていだだきます。\n",
}

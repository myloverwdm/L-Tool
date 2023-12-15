// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {index} from '../models';
import {configuration} from '../models';
import {onlineTools} from '../models';
import {goFunc} from '../models';

export function AddPath(arg1:string,arg2:string,arg3:string):Promise<string>;

export function DeleteZk(arg1:string):Promise<string>;

export function GetCommand():Promise<Array<index.CommandPojo>>;

export function GetItemList():Promise<Array<configuration.Index>>;

export function GetKafkaForm():Promise<index.KafkaForm>;

export function GetLanguageList():Promise<Array<configuration.Language>>;

export function GetNowLanguage():Promise<string>;

export function GetOnlineToolsList():Promise<Array<onlineTools.OnlineTool>>;

export function GetTopicAndPartition():Promise<Array<index.TopicAndPartition>>;

export function InitZk():Promise<index.ZkForm>;

export function JsonToXml(arg1:string):Promise<goFunc.LTResponse>;

export function ReadZkNode(arg1:string):Promise<index.ZkData>;

export function SetForm(arg1:index.ZkForm):Promise<string>;

export function SetNowLanguage(arg1:string):Promise<void>;

export function UpdateKafkaForm(arg1:index.KafkaForm):Promise<string>;
export namespace configuration {
	
	export class Index {
	    name: string;
	    route: string;
	    iconClass: string;
	
	    static createFrom(source: any = {}) {
	        return new Index(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.route = source["route"];
	        this.iconClass = source["iconClass"];
	    }
	}
	export class Language {
	    name: string;
	    languageCode: string;
	
	    static createFrom(source: any = {}) {
	        return new Language(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.languageCode = source["languageCode"];
	    }
	}

}

export namespace goFunc {
	
	export class LTResponse {
	    success: boolean;
	    data: string;
	    errorMsg: string;
	
	    static createFrom(source: any = {}) {
	        return new LTResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.data = source["data"];
	        this.errorMsg = source["errorMsg"];
	    }
	}

}

export namespace index {
	
	export class OneCommand {
	    desc: string;
	    command: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new OneCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.desc = source["desc"];
	        this.command = source["command"];
	        this.url = source["url"];
	    }
	}
	export class CommandPojo {
	    name: string;
	    commands: OneCommand[];
	
	    static createFrom(source: any = {}) {
	        return new CommandPojo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.commands = this.convertValues(source["commands"], OneCommand);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class KafkaForm {
	    connection: string;
	
	    static createFrom(source: any = {}) {
	        return new KafkaForm(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.connection = source["connection"];
	    }
	}
	
	export class Stat {
	    czxid?: number;
	    mzxid?: number;
	    ctime?: number;
	    mtime?: number;
	    version?: number;
	    cversion?: number;
	    aversion?: number;
	    ephemeralOwner?: number;
	    dataLength?: number;
	    numChildren?: number;
	    pzxid?: number;
	
	    static createFrom(source: any = {}) {
	        return new Stat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.czxid = source["czxid"];
	        this.mzxid = source["mzxid"];
	        this.ctime = source["ctime"];
	        this.mtime = source["mtime"];
	        this.version = source["version"];
	        this.cversion = source["cversion"];
	        this.aversion = source["aversion"];
	        this.ephemeralOwner = source["ephemeralOwner"];
	        this.dataLength = source["dataLength"];
	        this.numChildren = source["numChildren"];
	        this.pzxid = source["pzxid"];
	    }
	}
	export class TopicAndPartition {
	    topicName: string;
	    partitions: number[];
	
	    static createFrom(source: any = {}) {
	        return new TopicAndPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = source["partitions"];
	    }
	}
	export class ZkData {
	    nextPath: string[];
	    data: string;
	    statData: Stat;
	
	    static createFrom(source: any = {}) {
	        return new ZkData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nextPath = source["nextPath"];
	        this.data = source["data"];
	        this.statData = this.convertValues(source["statData"], Stat);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ZkForm {
	    address: string;
	    zkData: ZkData;
	
	    static createFrom(source: any = {}) {
	        return new ZkForm(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = source["address"];
	        this.zkData = this.convertValues(source["zkData"], ZkData);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace onlineTools {
	
	export class OnlineTool {
	    name: string;
	    desc: string;
	    router: string;
	
	    static createFrom(source: any = {}) {
	        return new OnlineTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.desc = source["desc"];
	        this.router = source["router"];
	    }
	}

}


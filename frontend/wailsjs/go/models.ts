export namespace backend {
	
	export class HotKey {
	    id: string;
	    key: string;
	    modifiers: string[];
	    width: number;
	    height: number;
	
	    static createFrom(source: any = {}) {
	        return new HotKey(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.key = source["key"];
	        this.modifiers = source["modifiers"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}
	export class Key {
	    id: string;
	
	    static createFrom(source: any = {}) {
	        return new Key(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}

}


export namespace config {
	
	export class OpenAI {
	    api_key: string;
	    base_url: string;
	    api_version: string;
	    model: string;
	
	    static createFrom(source: any = {}) {
	        return new OpenAI(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.api_key = source["api_key"];
	        this.base_url = source["base_url"];
	        this.api_version = source["api_version"];
	        this.model = source["model"];
	    }
	}
	export class GuiConfig {
	    openai: OpenAI;
	
	    static createFrom(source: any = {}) {
	        return new GuiConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.openai = this.convertValues(source["openai"], OpenAI);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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


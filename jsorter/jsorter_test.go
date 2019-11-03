package jsorter

import (
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	ts := []struct {
		desc        string
		orig        string
		reverse     bool
		errExpected bool
		exp         string
	}{
		{
			desc:        "Invalid JSON 1",
			orig:        "I'm JSON",
			errExpected: true,
		},
		{
			desc:        "Invalid JSON 2",
			orig:        `{"example":2:]}}`,
			errExpected: true,
		},
		{
			desc: "Valid JSON 1",
			orig: `{
  "name":"John",
  "age":30
}`,
			exp: `{
  "age":30,
  "name": "John"
}`,
		},
		// https://json.org/example.html
		{
			desc: "Valid JSON 2",
			orig: `{
  "glossary":{
    "title":"example glossary",
    "GlossDiv":{
      "title":"S",
      "GlossList":{
        "GlossEntry":{
          "ID":"SGML",
          "SortAs":"SGML",
          "GlossTerm":"Standard Generalized Markup Language",
          "Acronym":"SGML",
          "Abbrev":"ISO 8879:1986",
          "GlossDef":{
            "para":"A meta-markup language, used to create markup languages such as DocBook.",
            "GlossSeeAlso":[
              "GML",
              "XML"
            ]
          },
          "GlossSee":"markup"
        }
      }
    }
  }
}`,
			exp: `{
  "glossary":{
    "GlossDiv":{
      "GlossList":{
        "GlossEntry":{
          "Abbrev":"ISO 8879:1986",
          "Acronym":"SGML",
          "GlossDef":{
            "GlossSeeAlso":[
              "GML",
              "XML"
            ],
            "para":"A meta-markup language, used to create markup languages such as DocBook."
          },
          "GlossSee":"markup",
          "GlossTerm":"Standard Generalized Markup Language",
          "ID":"SGML",
          "SortAs":"SGML"
        }
      },
      "title":"S"
    },
    "title":"example glossary"
  }
}`,
		},
		{
			desc: "Valid JSON 3",
			orig: `{
  "menu":{
    "id":"file",
    "value":"File",
    "popup":{
      "menuitem":[
        {
          "value":"New",
          "onclick":"CreateNewDoc()"
        },
        {
          "value":"Open",
          "onclick":"OpenDoc()"
        },
        {
          "value":"Close",
          "onclick":"CloseDoc()"
        }
      ]
    }
  }
}`,
			exp: `{
  "menu":{
    "id":"file",
    "popup":{
      "menuitem":[
        {
          "onclick":"CreateNewDoc()",
          "value":"New"
        },
        {
          "onclick":"OpenDoc()",
          "value":"Open"
        },
        {
          "onclick":"CloseDoc()",
          "value":"Close"
        }
      ]
    },
    "value":"File"
  }
}`,
		},
		{
			desc: "Valid JSON 4",
			orig: `{
  "widget":{
    "debug":"on",
    "window":{
      "title":"Sample Konfabulator Widget",
      "name":"main_window",
      "width":500,
      "height":500
    },
    "image":{
      "src":"Images/Sun.png",
      "name":"sun1",
      "hOffset":250,
      "vOffset":250,
      "alignment":"center"
    },
    "text":{
      "data":"Click Here",
      "size":36,
      "style":"bold",
      "name":"text1",
      "hOffset":250,
      "vOffset":100,
      "alignment":"center",
      "onMouseUp":"sun1.opacity = (sun1.opacity / 100) * 90;"
    }
  }
}`,
			exp: `{
  "widget":{
    "debug":"on",
    "image":{
      "alignment":"center",
      "hOffset":250,
      "name":"sun1",
      "src":"Images/Sun.png",
      "vOffset":250
    },
    "text":{
      "alignment":"center",
      "data":"Click Here",
      "hOffset":250,
      "name":"text1",
      "onMouseUp":"sun1.opacity = (sun1.opacity / 100) * 90;",
      "size":36,
      "style":"bold",
      "vOffset":100
    },
    "window":{
      "height":500,
      "name":"main_window",
      "title":"Sample Konfabulator Widget",
      "width":500
    }
  }
}`,
		},
		{
			desc: "Valid JSON 5",
			orig: `{
  "web-app":{
    "servlet":[
      {
        "servlet-name":"cofaxCDS",
        "servlet-class":"org.cofax.cds.CDSServlet",
        "init-param":{
          "configGlossary:installationAt":"Philadelphia, PA",
          "configGlossary:adminEmail":"ksm@pobox.com",
          "configGlossary:poweredBy":"Cofax",
          "configGlossary:poweredByIcon":"/images/cofax.gif",
          "configGlossary:staticPath":"/content/static",
          "templateProcessorClass":"org.cofax.WysiwygTemplate",
          "templateLoaderClass":"org.cofax.FilesTemplateLoader",
          "templatePath":"templates",
          "templateOverridePath":"",
          "defaultListTemplate":"listTemplate.htm",
          "defaultFileTemplate":"articleTemplate.htm",
          "useJSP":false,
          "jspListTemplate":"listTemplate.jsp",
          "jspFileTemplate":"articleTemplate.jsp",
          "cachePackageTagsTrack":200,
          "cachePackageTagsStore":200,
          "cachePackageTagsRefresh":60,
          "cacheTemplatesTrack":100,
          "cacheTemplatesStore":50,
          "cacheTemplatesRefresh":15,
          "cachePagesTrack":200,
          "cachePagesStore":100,
          "cachePagesRefresh":10,
          "cachePagesDirtyRead":10,
          "searchEngineListTemplate":"forSearchEnginesList.htm",
          "searchEngineFileTemplate":"forSearchEngines.htm",
          "searchEngineRobotsDb":"WEB-INF/robots.db",
          "useDataStore":true,
          "dataStoreClass":"org.cofax.SqlDataStore",
          "redirectionClass":"org.cofax.SqlRedirection",
          "dataStoreName":"cofax",
          "dataStoreDriver":"com.microsoft.jdbc.sqlserver.SQLServerDriver",
          "dataStoreUrl":"jdbc:microsoft:sqlserver://LOCALHOST:1433;DatabaseName=goon",
          "dataStoreUser":"sa",
          "dataStorePassword":"dataStoreTestQuery",
          "dataStoreTestQuery":"SET NOCOUNT ON;select test='test';",
          "dataStoreLogFile":"/usr/local/tomcat/logs/datastore.log",
          "dataStoreInitConns":10,
          "dataStoreMaxConns":100,
          "dataStoreConnUsageLimit":100,
          "dataStoreLogLevel":"debug",
          "maxUrlLength":500
        }
      },
      {
        "servlet-name":"cofaxEmail",
        "servlet-class":"org.cofax.cds.EmailServlet",
        "init-param":{
          "mailHost":"mail1",
          "mailHostOverride":"mail2"
        }
      },
      {
        "servlet-name":"cofaxAdmin",
        "servlet-class":"org.cofax.cds.AdminServlet"
      },
      {
        "servlet-name":"fileServlet",
        "servlet-class":"org.cofax.cds.FileServlet"
      },
      {
        "servlet-name":"cofaxTools",
        "servlet-class":"org.cofax.cms.CofaxToolsServlet",
        "init-param":{
          "templatePath":"toolstemplates/",
          "log":1,
          "logLocation":"/usr/local/tomcat/logs/CofaxTools.log",
          "logMaxSize":"",
          "dataLog":1,
          "dataLogLocation":"/usr/local/tomcat/logs/dataLog.log",
          "dataLogMaxSize":"",
          "removePageCache":"/content/admin/remove?cache=pages&id=",
          "removeTemplateCache":"/content/admin/remove?cache=templates&id=",
          "fileTransferFolder":"/usr/local/tomcat/webapps/content/fileTransferFolder",
          "lookInContext":1,
          "adminGroupID":4,
          "betaServer":true
        }
      }
    ],
    "servlet-mapping":{
      "cofaxCDS":"/",
      "cofaxEmail":"/cofaxutil/aemail/*",
      "cofaxAdmin":"/admin/*",
      "fileServlet":"/static/*",
      "cofaxTools":"/tools/*"
    },
    "taglib":{
      "taglib-uri":"cofax.tld",
      "taglib-location":"/WEB-INF/tlds/cofax.tld"
    }
  }
}`,
			exp: `{
  "web-app": {
    "servlet": [
      {
        "init-param": {
          "cachePackageTagsRefresh": 60,
          "cachePackageTagsStore": 200,
          "cachePackageTagsTrack": 200,
          "cachePagesDirtyRead": 10,
          "cachePagesRefresh": 10,
          "cachePagesStore": 100,
          "cachePagesTrack": 200,
          "cacheTemplatesRefresh": 15,
          "cacheTemplatesStore": 50,
          "cacheTemplatesTrack": 100,
          "configGlossary:adminEmail": "ksm@pobox.com",
          "configGlossary:installationAt": "Philadelphia, PA",
          "configGlossary:poweredBy": "Cofax",
          "configGlossary:poweredByIcon": "/images/cofax.gif",
          "configGlossary:staticPath": "/content/static",
          "dataStoreClass": "org.cofax.SqlDataStore",
          "dataStoreConnUsageLimit": 100,
          "dataStoreDriver": "com.microsoft.jdbc.sqlserver.SQLServerDriver",
          "dataStoreInitConns": 10,
          "dataStoreLogFile": "/usr/local/tomcat/logs/datastore.log",
          "dataStoreLogLevel": "debug",
          "dataStoreMaxConns": 100,
          "dataStoreName": "cofax",
          "dataStorePassword": "dataStoreTestQuery",
          "dataStoreTestQuery": "SET NOCOUNT ON;select test='test';",
          "dataStoreUrl": "jdbc:microsoft:sqlserver://LOCALHOST:1433;DatabaseName=goon",
          "dataStoreUser": "sa",
          "defaultFileTemplate": "articleTemplate.htm",
          "defaultListTemplate": "listTemplate.htm",
          "jspFileTemplate": "articleTemplate.jsp",
          "jspListTemplate": "listTemplate.jsp",
          "maxUrlLength": 500,
          "redirectionClass": "org.cofax.SqlRedirection",
          "searchEngineFileTemplate": "forSearchEngines.htm",
          "searchEngineListTemplate": "forSearchEnginesList.htm",
          "searchEngineRobotsDb": "WEB-INF/robots.db",
          "templateLoaderClass": "org.cofax.FilesTemplateLoader",
          "templateOverridePath": "",
          "templatePath": "templates",
          "templateProcessorClass": "org.cofax.WysiwygTemplate",
          "useDataStore": true,
          "useJSP": false
        },
        "servlet-class": "org.cofax.cds.CDSServlet",
        "servlet-name": "cofaxCDS"
      },
      {
        "init-param": {
          "mailHost": "mail1",
          "mailHostOverride": "mail2"
        },
        "servlet-class": "org.cofax.cds.EmailServlet",
        "servlet-name": "cofaxEmail"
      },
      {
        "servlet-class": "org.cofax.cds.AdminServlet",
        "servlet-name": "cofaxAdmin"
      },
      {
        "servlet-class": "org.cofax.cds.FileServlet",
        "servlet-name": "fileServlet"
      },
      {
        "init-param": {
          "adminGroupID": 4,
          "betaServer": true,
          "dataLog": 1,
          "dataLogLocation": "/usr/local/tomcat/logs/dataLog.log",
          "dataLogMaxSize": "",
          "fileTransferFolder": "/usr/local/tomcat/webapps/content/fileTransferFolder",
          "log": 1,
          "logLocation": "/usr/local/tomcat/logs/CofaxTools.log",
          "logMaxSize": "",
          "lookInContext": 1,
          "removePageCache": "/content/admin/remove?cache=pages\u0026id=",
          "removeTemplateCache": "/content/admin/remove?cache=templates\u0026id=",
          "templatePath": "toolstemplates/"
        },
        "servlet-class": "org.cofax.cms.CofaxToolsServlet",
        "servlet-name": "cofaxTools"
      }
    ],
    "servlet-mapping": {
      "cofaxAdmin": "/admin/*",
      "cofaxCDS": "/",
      "cofaxEmail": "/cofaxutil/aemail/*",
      "cofaxTools": "/tools/*",
      "fileServlet": "/static/*"
    },
    "taglib": {
      "taglib-location": "/WEB-INF/tlds/cofax.tld",
      "taglib-uri": "cofax.tld"
    }
  }
}`,
		},
	}

	for _, tc := range ts {
		t.Log(tc.desc)

		result, err := Sort([]byte(tc.orig), tc.reverse)

		if tc.errExpected {
			if err == nil {
				t.Error("Expected error, got none.")
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		if strings.ReplaceAll(string(result), " ", "") != strings.ReplaceAll(tc.exp, " ", "") {
			t.Error("Not sorted as expected!")
			t.Error("Original:", tc.orig)
			t.Error("Expected:", tc.exp)
			t.Error("Got:", string(result))
			t.FailNow()
		}
	}
}

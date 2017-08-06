# testcase
Small Golang util to prepare few test cases inside single test & run them. Usage example: 

```go
cs := testcase.Cases{
    {
        Name:   "test_case_name (like 'create backup of non-existing file')",
        Params: testcase.Params{
            //BeforeArgs - arguments for tcase.Before() callback 
            BeforeArgs: map[string]interface{}{},
            //AfterArgs - arguments for tcase.After() callback
            AfterArgs: map[string]interface{}{},
            //PassedArgs - arguments for tcase.Passed() callback 
            PassedArgs: map[string]interface{} {
                "file":   "file.tmp",
            },
        },
        Before: func(args map[string]interface{}) {
            //execute some actions before test (create file / check that folder empty / etc)
        },
        After:  func(args map[string]interface{}) {
            //execute some actions after test (cleanup / etc)
        },
        Passed: func(args map[string]interface{}) bool {
            //test body here. smth like: 
            return args["backup"].(*Backup).Create(args["file"].(string)) != nil
        },
    },
}

testcase.Run(t, cs)
```
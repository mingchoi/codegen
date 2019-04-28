{{if .Equal "ClassName" "MyClass"}}// MyClass{{end}}
{{if .Equal "ClassName" "YourClass"}}// YourClass{{end}}
public class {{.Print "ClassName"}} {
    private static final String {{.Print "MyVariable"}} = "HelloWorld";
    
    public static void main(String args[]){
        System.out.println({{.Print "MyVariable"}});
    }
}
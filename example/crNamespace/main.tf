resource "volcengine_cr_namespace" "foo"{
     registry = "tf-2"
     name ="namespace-1"
}

resource "volcengine_cr_namespace" "foo1"{
     registry = "tf-1"
     name ="namespace-2"
}
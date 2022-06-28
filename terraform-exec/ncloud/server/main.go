package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func main() {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
		return
	}

	//workingDir := "C:\\Users\\kng\\Documents\\GitHub\\nk915\\terraform_example\\terrform\\ncloud\\server"
	workingDir := "..\\..\\..\\terrform\\ncloud\\server"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
		return
	}

	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
		return
	}
	result, err := tf.Plan(context.Background())
	if err != nil {
		log.Fatalf("error running Plan: %s", err)
		return
	}
	if result == false {
		fmt.Println("tf plan: ", result)
		return
	}

	//	if err := tf.Destroy(context.Background()); err != nil {
	//		fmt.Println("tf destroy success")
	//		return
	//	}
	//	fmt.Println("tf destroy success")
	//	return

	if err := tf.Apply(context.Background()); err != nil {
		log.Fatalf("error running Apply: %s", err)
		return
	}
	fmt.Println("tf apply success")

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
		return
	}

	fmt.Println(state) // "0.1"
}

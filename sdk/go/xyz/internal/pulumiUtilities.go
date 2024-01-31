// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package internal

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type envParser func(v string) interface{}

func ParseEnvBool(v string) interface{} {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return nil
	}
	return b
}

func ParseEnvInt(v string) interface{} {
	i, err := strconv.ParseInt(v, 0, 0)
	if err != nil {
		return nil
	}
	return int(i)
}

func ParseEnvFloat(v string) interface{} {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil
	}
	return f
}

func ParseEnvStringArray(v string) interface{} {
	var result pulumi.StringArray
	for _, item := range strings.Split(v, ";") {
		result = append(result, pulumi.String(item))
	}
	return result
}

func GetEnvOrDefault(def interface{}, parser envParser, vars ...string) interface{} {
	for _, v := range vars {
		if value, ok := os.LookupEnv(v); ok {
			if parser != nil {
				return parser(value)
			}
			return value
		}
	}
	return def
}

// PkgVersion uses reflection to determine the version of the current package.
// If a version cannot be determined, v1 will be assumed. The second return
// value is always nil.
func PkgVersion() (semver.Version, error) {
	// emptyVersion defaults to v0.0.0
	if !SdkVersion.Equals(semver.Version{}) {
		return SdkVersion, nil
	}
	type sentinal struct{}
	pkgPath := reflect.TypeOf(sentinal{}).PkgPath()
	re := regexp.MustCompile("^.*/pulumi-asd/sdk(/v\\d+)?")
	if match := re.FindStringSubmatch(pkgPath); match != nil {
		vStr := match[1]
		if len(vStr) == 0 { // If the version capture group was empty, default to v1.
			return semver.Version{Major: 1}, nil
		}
		return semver.MustParse(fmt.Sprintf("%s.0.0", vStr[2:])), nil
	}
	return semver.Version{Major: 1}, nil
}

// isZero is a null safe check for if a value is it's types zero value.
func IsZero(v interface{}) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsZero()
}

// PkgResourceDefaultOpts provides package level defaults to pulumi.OptionResource.
func PkgResourceDefaultOpts(opts []pulumi.ResourceOption) []pulumi.ResourceOption {
	defaults := []pulumi.ResourceOption{}

	version := SdkVersion
	if !version.Equals(semver.Version{}) {
		defaults = append(defaults, pulumi.Version(version.String()))
	}
	return append(defaults, opts...)
}

// PkgInvokeDefaultOpts provides package level defaults to pulumi.OptionInvoke.
func PkgInvokeDefaultOpts(opts []pulumi.InvokeOption) []pulumi.InvokeOption {
	defaults := []pulumi.InvokeOption{}

	version := SdkVersion
	if !version.Equals(semver.Version{}) {
		defaults = append(defaults, pulumi.Version(version.String()))
	}
	return append(defaults, opts...)
}

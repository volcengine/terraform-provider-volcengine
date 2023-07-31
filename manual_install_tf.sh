#!/bin/bash

_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
_FILE="${_DIR}/$(basename "${BASH_SOURCE[0]}")"
_BASE="$(basename "${_FILE}")"

INSTALLATION_DIR="${HOME}/.terraform.d/plugins"

# usage prints script usage
function usage(){
    echo ""
    echo "Usage: ${_BASE} [options]"
    echo " "
    echo "Prerequisite: you must install terraform, wget, xattr at first"
    echo " "
    echo "Installs the Volcengine Terraform Provider"
    echo "- For Terraform 0.12, the installation path will be: ${INSTALLATION_DIR}"
    echo "- For Terraform > 0.12, the installation path will be: ${INSTALLATION_DIR}/volcengine/volcengine/PROVIDER_VERSION/OS_ARCH"
    echo " "
    echo "options:"
    echo "-h, --help                                            show help"
    echo "-v, --version                                         [required] specify the version of the terraform-provider-volcengine, e.g. 0.0.89"
    echo "-p, --provider                                        [optional] specify the local path of the provider zip file, if no value is specified, the file will be downloaded from github by default"
    echo "-a, --arch                                            [optional] specify the architecture, if no value is specified, the default value is amd64"
}

function determineOSAndArch(){
  # used to determine which os and architecture to install
  XC_OS=$(uname)
  XC_ARCH=${XC_ARCH:-"amd64"}

  # determine which binary to fetch depending upon current architecture
  if [ "${XC_OS}" == "Linux" ]; then
    XC_OS="linux"
  elif [ "${XC_OS}" == "Darwin" ]; then
    XC_OS="darwin"
  else
    echo "Unsupported architecture: ${XC_OS}, only architecture supported at the moment are Darwin and Linux"
    exit 2
  fi
}


# process input arguments
while [ $# -gt 0 ]; do
    case $1 in
        --help | -h)
            usage
            exit 0
            ;;
        --version | -v)
            shift
            PROVIDER_VERSION=$1
            ;;
        --provider | -p)
            shift
            PROVIDER=$1
            ;;
        --arch | -a)
            shift
            XC_ARCH=$1
            ;;
        *)
            echo "Unknown option: $1"
            usage
            exit 2
    esac
    shift
done

if [ "$PROVIDER_VERSION" == "" ]; then
  echo "required argument --version missing value."
  usage
  exit 1
fi

determineOSAndArch


TF_PATH=`pwd`

TF_PATH="$TF_PATH"/tf_install_template
if [ "$PROVIDER" == "" ]
then
   wget https://github.com/volcengine/terraform-provider-volcengine/releases/download/v"$PROVIDER_VERSION"/terraform-provider-volcengine_"$PROVIDER_VERSION"_"$XC_OS"_"$XC_ARCH".zip
   unzip terraform-provider-volcengine_"$PROVIDER_VERSION"_"$XC_OS"_"$XC_ARCH".zip -d "$TF_PATH"
else
   unzip "$PROVIDER" -d "$TF_PATH"
fi

PROVIDER_NAME="$TF_PATH"/terraform-provider-volcengine_v"$PROVIDER_VERSION"

PERMISSION=`xattr -l $PROVIDER_NAME`
if [ "$PERMISSION" != "" ]
then
  NUM=0
  for LINE in $PERMISSION;
  do
    if ((NUM % 2 == 0))
    then
      KEY=`echo "${LINE%:*}"`
      `xattr -d $KEY $PROVIDER_NAME`
    fi
    NUM=$NUM+1
  done
fi


TF_VERSION=`terraform -version | head -n 1`
TF_VERSION=`echo ${TF_VERSION#*v}`
if [ x"$TF_VERSION" \< x"0.13" ];then
  mkdir -p "$INSTALLATION_DIR"
  mv "$TF_PATH"/terraform-provider-volcengine_v"$PROVIDER_VERSION" "$INSTALLATION_DIR"/terraform-provider-volcengine_v"$PROVIDER_VERSION"
else
	mkdir -p "$INSTALLATION_DIR"/registry.terraform.io/volcengine/volcengine/"$PROVIDER_VERSION"/"$XC_OS"_"$XC_ARCH"/
  mv "$TF_PATH"/terraform-provider-volcengine_v"$PROVIDER_VERSION" "$INSTALLATION_DIR"/registry.terraform.io/volcengine/volcengine/"$PROVIDER_VERSION"/"$XC_OS"_"$XC_ARCH"/terraform-provider-volcengine_v"$PROVIDER_VERSION"
fi
rm -rf terraform-provider-volcengine*
rm -rf tf_install_template

echo "install completed"

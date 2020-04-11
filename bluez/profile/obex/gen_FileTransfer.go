// Code generated DO NOT EDIT

package obex



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var FileTransferInterface = "org.bluez.obex.FileTransfer"


// NewFileTransfer create a new instance of FileTransfer
//
// Args:
// - objectPath: [Session object path]
func NewFileTransfer(objectPath dbus.ObjectPath) (*FileTransfer, error) {
	a := new(FileTransfer)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez.obex",
			Iface: FileTransferInterface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(FileTransferProperties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
FileTransfer File Transfer hierarchy

*/
type FileTransfer struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*FileTransferProperties
	watchPropertiesChannel chan *dbus.Signal
}

// FileTransferProperties contains the exposed properties of an interface
type FileTransferProperties struct {
	lock sync.RWMutex `dbus:"ignore"`

}

//Lock access to properties
func (p *FileTransferProperties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *FileTransferProperties) Unlock() {
	p.lock.Unlock()
}



// Close the connection
func (a *FileTransfer) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return FileTransfer object path
func (a *FileTransfer) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return FileTransfer dbus client
func (a *FileTransfer) Client() *bluez.Client {
	return a.client
}

// Interface return FileTransfer interface
func (a *FileTransfer) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *FileTransfer) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a FileTransferProperties to map
func (a *FileTransferProperties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an FileTransferProperties
func (a *FileTransferProperties) FromMap(props map[string]interface{}) (*FileTransferProperties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an FileTransferProperties
func (a *FileTransferProperties) FromDBusMap(props map[string]dbus.Variant) (*FileTransferProperties, error) {
	s := new(FileTransferProperties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *FileTransfer) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *FileTransfer) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *FileTransfer) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *FileTransfer) GetProperties() (*FileTransferProperties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *FileTransfer) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *FileTransfer) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *FileTransfer) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *FileTransfer) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *FileTransfer) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *FileTransfer) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
ChangeFolder 
			Change the current folder of the remote device.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) ChangeFolder(folder string) error {
	
	return a.client.Call("ChangeFolder", 0, folder).Store()
	
}

/*
CreateFolder 
			Create a new folder in the remote device.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) CreateFolder(folder string) error {
	
	return a.client.Call("CreateFolder", 0, folder).Store()
	
}

/*
ListFolder 
			Returns a dictionary containing information about
			the current folder content.

			The following keys are defined:

				string Name : Object name in UTF-8 format
				string Type : Either "folder" or "file"
				uint64 Size : Object size or number of items in
						folder
				string Permission : Group, owner and other
							permission
				uint64 Modified : Last change
				uint64 Accessed : Last access
				uint64 Created : Creation date

			Possible errors: org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) ListFolder() ([]map[string]interface{}, error) {
	
	var val0 []map[string]interface{}
	err := a.client.Call("ListFolder", 0, ).Store(&val0)
	return val0, err	
}

/*
GetFile 
			Copy the source file (from remote device) to the
			target file (on local filesystem).

			If an empty target file is given, a name will be
			automatically calculated for the temporary file.

			The returned path represents the newly created transfer,
			which should be used to find out if the content has been
			successfully transferred or if the operation fails.

			The properties of this transfer are also returned along
			with the object path, to avoid a call to GetProperties.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) GetFile(targetfile string, sourcefile string) (dbus.ObjectPath, map[string]interface{}, error) {
	
	var val0 dbus.ObjectPath
  var val1 map[string]interface{}
	err := a.client.Call("GetFile", 0, targetfile, sourcefile).Store(&val0, &val1)
	return val0, val1, err	
}

/*
PutFile 
			Copy the source file (from local filesystem) to the
			target file (on remote device).

			The returned path represents the newly created transfer,
			which should be used to find out if the content has been
			successfully transferred or if the operation fails.

			The properties of this transfer are also returned along
			with the object path, to avoid a call to GetProperties.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) PutFile(sourcefile string, targetfile string) (dbus.ObjectPath, map[string]interface{}, error) {
	
	var val0 dbus.ObjectPath
  var val1 map[string]interface{}
	err := a.client.Call("PutFile", 0, sourcefile, targetfile).Store(&val0, &val1)
	return val0, val1, err	
}

/*
CopyFile 
			Copy a file within the remote device from source file
			to target file.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) CopyFile(sourcefile string, targetfile string) error {
	
	return a.client.Call("CopyFile", 0, sourcefile, targetfile).Store()
	
}

/*
MoveFile 
			Move a file within the remote device from source file
			to the target file.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed


*/
func (a *FileTransfer) MoveFile(sourcefile string, targetfile string) error {
	
	return a.client.Call("MoveFile", 0, sourcefile, targetfile).Store()
	
}

/*
Delete 
			Deletes the specified file/folder.

			Possible errors: org.bluez.obex.Error.InvalidArguments
					 org.bluez.obex.Error.Failed



*/
func (a *FileTransfer) Delete(file string) error {
	
	return a.client.Call("Delete", 0, file).Store()
	
}

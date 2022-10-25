package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/sparkkoori/go-vulkan/v1.1/vk"
)

// #cgo LDFLAGS: -L/Users/phil/VulkanSDK/1.2.182.0/macOS/lib/
import "C"

func init() {
	runtime.LockOSThread()
}

func main() {

	win, err := newWindow(640, 480, "Hello World")
	if err != nil {
		panic(err)
	}
	defer win.destroy()

	dev, err := newDevice(win)
	if err != nil {
		panic(err)
	}
	defer dev.destroy()

	glfw.SetMonitorCallback(func(monitor *glfw.Monitor, event glfw.PeripheralEvent) {
		mX, mY := monitor.GetPos()
		mW, mH := monitor.GetPhysicalSize()
		fmt.Printf("Monitor: %v %v %v %v %v\n", monitor.GetName(), mX, mY, mW, mH)
	})

	glfw.SetJoystickCallback(func(joy glfw.Joystick, event glfw.PeripheralEvent) {
		fmt.Printf("Joy stick: %v %v\n", joy, event)
	})

	for !win.shouldClose() {
		glfw.PollEvents()

		time.Sleep(20 * time.Millisecond)
	}

	vk.DeviceWaitIdle(dev.device)
}

type SwapChainSupportDetails struct {
	capabilities vk.SurfaceCapabilitiesKHR
	formats      []vk.SurfaceFormatKHR
	presentModes []vk.PresentModeKHR
}

type QueueFamilyIndices struct {
	graphicsFamily         uint32
	presentFamily          uint32
	graphicsFamilyHasValue bool
	presentFamilyHasValue  bool
}

func (this *QueueFamilyIndices) isComplete() bool {
	return this.graphicsFamilyHasValue && this.presentFamilyHasValue
}

type AxeDevice struct {
	properties                  vk.PhysicalDeviceProperties
	instance                    vk.Instance
	instanceAvailableExtensions instanceExtensions
	instanceEnabledExtensions   instanceExtensions
	debugMessenger              vk.DebugUtilsMessengerEXT
	physicalDevice              vk.PhysicalDevice
	window                      *AxeWindow
	commandPool                 vk.CommandPool
	device                      vk.Device
	deviceActualExtensions      deviceExtensions
	surface                     vk.SurfaceKHR
	graphicsQueue               vk.Queue
	presentQueue                vk.Queue
	validationLayers            layerExtensions
	deviceEnabledExtensions     deviceExtensions
	deviceRequiredExtensions    deviceExtensions
	enableValidationLayers      bool
}

type instanceExtensions struct {
	DeviceGroupCreation           bool `name:"VK_KHR_device_group_creation"`
	ExternalFenceCapabilities     bool `name:"VK_KHR_external_fence_capabilities"`
	ExternalMemoryCapabilities    bool `name:"VK_KHR_external_memory_capabilities"`
	ExternalSemaphoreCapabilities bool `name:"VK_KHR_external_semaphore_capabilities"`
	PhysicalDeviceProperties      bool `name:"VK_KHR_get_physical_device_properties2"`
	SurfaceCapabilities           bool `name:"VK_KHR_get_surface_capabilities2"`
	Surface                       bool `name:"VK_KHR_surface"`
	DebugReport                   bool `name:"VK_EXT_debug_report"`
	DebugUtils                    bool `name:"VK_EXT_debug_utils"`
	MetalSurface                  bool `name:"VK_EXT_metal_surface"`
	SwapchainColorspace           bool `name:"VK_EXT_swapchain_colorspace"`
	MacOsSurface                  bool `name:"VK_MVK_macos_surface"`
}

func (ext *instanceExtensions) fromNames(names []string) { setFromNames(ext, names) }
func (ext *instanceExtensions) fromLayerName(layerName string) error {
	extenstionProperties, err := getInstanceExtensionProperties(layerName)
	if err != nil {
		return err
	}

	ext.fromNames(MapSlice(extenstionProperties, func(ext vk.ExtensionProperties) string {
		return strings.Trim(string(ext.ExtensionName[:]), "\x00")
	}))
	return nil
}
func (ext instanceExtensions) isSupported(expected instanceExtensions) bool {
	return len(ext.getMissing(expected)) == 0
}
func (ext instanceExtensions) getMissing(expected instanceExtensions) []string {
	return getMissingNames(ext, expected)
}
func (ext instanceExtensions) getExtensions() []string { return instanceExtensions{}.getMissing(ext) }

type layerExtensions struct {
	ApiDump          bool `name:"VK_LAYER_LUNARG_api_dump"`
	Validation       bool `name:"VK_LAYER_KHRONOS_validation"`
	DeviceSimulation bool `name:"VK_LAYER_LUNARG_device_simulation"`
	Synchronization  bool `name:"VK_LAYER_KHRONOS_synchronization2"`
}

func (ext *layerExtensions) fromNames(names []string) { setFromNames(ext, names) }
func (ext *layerExtensions) fromSystem() error {
	layerProperties, err := getInstanceLayerProperties()
	if err != nil {
		return err
	}
	ext.fromNames(MapSlice(layerProperties, func(source vk.LayerProperties) string {
		return strings.Trim(string(source.LayerName[:]), "\x00")
	}))
	return err
}
func (ext layerExtensions) isSupported(expected layerExtensions) bool {
	return len(ext.getMissing(expected)) == 0
}
func (ext layerExtensions) getMissing(expected layerExtensions) []string {
	return getMissingNames(ext, expected)
}
func (ext layerExtensions) getExtensions() []string { return layerExtensions{}.getMissing(ext) }

type deviceExtensions struct {
	GpuShaderHalfFloat          bool `name:"VK_AMD_gpu_shader_half_float"`
	NegativeViewportHeight      bool `name:"VK_AMD_negative_viewport_height"`
	DebugMark                   bool `name:"VK_EXT_debug_marker"`
	DescriptorIndexing          bool `name:"VK_EXT_descriptor_indexing"`
	FragmentShaderInterlock     bool `name:"VK_EXT_fragment_shader_interlock"`
	ImageRobustness             bool `name:"VK_EXT_image_robustness"`
	InlineUniformBlock          bool `name:"VK_EXT_inline_uniform_block"`
	MemoryBudget                bool `name:"VK_EXT_memory_budget"`
	PrivateData                 bool `name:"VK_EXT_private_data"`
	Robustness2                 bool `name:"VK_EXT_robustness2"`
	ScalarBlockLayout           bool `name:"VK_EXT_scalar_block_layout"`
	ShaderViewportIndexLayer    bool `name:"VK_EXT_shader_viewport_index_layer"`
	TexelBufferAlignment        bool `name:"VK_EXT_texel_buffer_alignment"`
	VertexAttributeDivisor      bool `name:"VK_EXT_vertex_attribute_divisor"`
	DisplayTiming               bool `name:"VK_GOOGLE_display_timing"`
	ShaderIntegerFunctions2     bool `name:"VK_INTEL_shader_integer_functions2"`
	Storage16bit                bool `name:"VK_KHR_16bit_storage"`
	Storage8bit                 bool `name:"VK_KHR_8bit_storage"`
	BindMemory2                 bool `name:"VK_KHR_bind_memory2"`
	CreateRenderpass2           bool `name:"VK_KHR_create_renderpass2"`
	DedicatedAllocation         bool `name:"VK_KHR_dedicated_allocation"`
	DepthStencilResolve         bool `name:"VK_KHR_depth_stencil_resolve"`
	DescriptorUpdateTemplate    bool `name:"VK_KHR_descriptor_update_template"`
	DeviceGroup                 bool `name:"VK_KHR_device_group"`
	DriverProperties            bool `name:"VK_KHR_driver_properties"`
	ExternalFence               bool `name:"VK_KHR_external_fence"`
	ExternalMemory              bool `name:"VK_KHR_external_memory"`
	ExternalSemaphore           bool `name:"VK_KHR_external_semaphore"`
	MemoryRequirements          bool `name:"VK_KHR_get_memory_requirements2"`
	ImageFormatList             bool `name:"VK_KHR_image_format_list"`
	ImagelessFramebuffer        bool `name:"VK_KHR_imageless_framebuffer"`
	Maintenance1                bool `name:"VK_KHR_maintenance1"`
	Maintenance2                bool `name:"VK_KHR_maintenance2"`
	Maintenance3                bool `name:"VK_KHR_maintenance3"`
	Multiview                   bool `name:"VK_KHR_multiview"`
	PortabilitySubset           bool `name:"VK_KHR_portability_subset"`
	PushDescriptor              bool `name:"VK_KHR_push_descriptor"`
	RelaxedBlockLayout          bool `name:"VK_KHR_relaxed_block_layout"`
	SamplerMirrorClampToEdge    bool `name:"VK_KHR_sampler_mirror_clamp_to_edge"`
	SamplerYcbcrConversion      bool `name:"VK_KHR_sampler_ycbcr_conversion"`
	ShaderDrawParameters        bool `name:"VK_KHR_shader_draw_parameters"`
	ShaderFloat16Int8           bool `name:"VK_KHR_shader_float16_int8"`
	StorageBufferStorageClass   bool `name:"VK_KHR_storage_buffer_storage_class"`
	Swapchain                   bool `name:"VK_KHR_swapchain"`
	SwapchainMutableFormat      bool `name:"VK_KHR_swapchain_mutable_format"`
	TimelineSemaphore           bool `name:"VK_KHR_timeline_semaphore"`
	UniformBufferStandardLayout bool `name:"VK_KHR_uniform_buffer_standard_layout"`
	VariablePointers            bool `name:"VK_KHR_variable_pointers"`
	GlslShader                  bool `name:"VK_NV_glsl_shader"`
}

func (ext *deviceExtensions) fromNames(names []string) { setFromNames(ext, names) }
func (ext *deviceExtensions) fromDevice(physicalDevice vk.PhysicalDevice, layerName string) error {
	extensionProperties, err := getDeviceExtensionProperties(physicalDevice, layerName)
	if err != nil {
		return err
	}
	ext.fromNames(MapSlice(extensionProperties, func(source vk.ExtensionProperties) string {
		return strings.Trim(string(source.ExtensionName[:]), "\x00")
	}))
	return nil
}

func (ext deviceExtensions) isSupported(expected deviceExtensions) bool {
	return len(ext.getMissing(expected)) == 0
}
func (ext deviceExtensions) getMissing(expected deviceExtensions) []string {
	return getMissingNames(ext, expected)
}
func (ext deviceExtensions) getExtensions() []string { return deviceExtensions{}.getMissing(ext) }

func newDevice(window *AxeWindow) (*AxeDevice, error) {
	dev := new(AxeDevice)
	dev.window = window
	dev.validationLayers.Validation = true
	dev.deviceRequiredExtensions.Swapchain = true
	dev.deviceEnabledExtensions.Swapchain = true
	dev.instanceEnabledExtensions.Surface = true
	// dev.instanceEnabledExtensions.SurfaceCapabilities = true
	dev.instanceEnabledExtensions.PhysicalDeviceProperties = true
	dev.enableValidationLayers = true

	err := dev.loadInstanceExtensions()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v]\n", dev.instanceAvailableExtensions)

	err = dev.createInstance()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Instance enabled extensions: %v\n", dev.getRequiredExtensions().getExtensions())

	dev.setupDebugMessenger()

	err = dev.createSurface()
	if err != nil {
		return nil, err
	}

	err = dev.pickPhysicalDevice()
	if err != nil {
		return nil, err
	}

	err = dev.createLogicalDevice()
	if err != nil {
		return nil, err
	}

	err = dev.createCommandPool()
	if err != nil {
		return nil, err
	}

	return dev, nil
}

func (this *AxeDevice) destroy() {
	if this == nil {
		return
	}
	if this.commandPool != 0 {
		vk.DestroyCommandPool(this.device, this.commandPool, nil)
	}
	vk.DestroyDevice(this.device, nil)
	if this.enableValidationLayers {
		vk.ToDestroyDebugUtilsMessengerEXT(
			vk.GetInstanceProcAddr(this.instance, "DestroyDebugUtilsMessengerEXT"),
		)(this.instance, this.debugMessenger, nil)
	}

	if this.surface != 0 {
		vk.ToDestroySurfaceKHR(
			vk.GetInstanceProcAddr(this.instance, "vkDestroySurfaceKHR"),
		)(this.instance, this.surface, nil)
	}

	vk.DestroyInstance(this.instance, nil)
}

func (this *AxeDevice) loadInstanceExtensions() error {
	return this.instanceAvailableExtensions.fromLayerName("")
}

func (this *AxeDevice) createInstance() error {
	if this.enableValidationLayers {
		supported, err := this.checkValidationLayerSupport()
		if !supported {
			if err != nil {
				return err
			} else {
				return errors.New("validation layers requested, but not available!")
			}
		}
	}

	appInfo := vk.ApplicationInfo{}
	appInfo.ApplicationName = "Axe App"
	appInfo.ApplicationVersion = vk.MAKE_VERSION(1, 0, 0)
	appInfo.EngineName = "Axe Game Engine"
	appInfo.EngineVersion = vk.MAKE_VERSION(1, 0, 0)
	appInfo.ApiVersion = vk.MAKE_VERSION(1, 1, 0)

	createInfo := vk.InstanceCreateInfo{}
	createInfo.ApplicationInfo = &appInfo

	extensions := this.getRequiredExtensions()
	createInfo.EnabledExtensionNames = extensions.getExtensions()
	// createInfo.Flags = vk.InstanceCreateFlags(vk.ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT)

	if this.enableValidationLayers {
		debugCreateInfo := vk.DebugUtilsMessengerCreateInfoEXT{}
		createInfo.EnabledLayerNames = this.validationLayers.getExtensions()

		this.populateDebugMessengerCreateInfo(&debugCreateInfo)
		createInfo.Next = &debugCreateInfo
	}

	if vk.CreateInstance(&createInfo, nil, &this.instance) != vk.SUCCESS {
		return errors.New("failed to create instance!")
	}

	err := this.hasGflwRequiredInstanceExtensions()
	if err != nil {
		return err
	}

	return nil
}
func (this *AxeDevice) setupDebugMessenger() error {
	if !this.enableValidationLayers {
		return nil
	}
	createInfo := vk.DebugUtilsMessengerCreateInfoEXT{}
	this.populateDebugMessengerCreateInfo(&createInfo)

	CreateDebugUtilsMessengerEXT := vk.GetInstanceProcAddr(this.instance, "CreateDebugUtilsMessengerEXT")
	if CreateDebugUtilsMessengerEXT != nil {
		if vk.ToCreateDebugUtilsMessengerEXT(CreateDebugUtilsMessengerEXT)(this.instance, &createInfo, nil, &this.debugMessenger) != vk.SUCCESS {
			return errors.New("failed to set up debug messenger!")
		}
	}
	return nil
}
func (this *AxeDevice) createSurface() error {
	surface, err := this.window.createWindowSurface(this.instance)
	this.surface = surface
	return err
}
func (this *AxeDevice) pickPhysicalDevice() error {
	devices, err := getPhysicalDevices(this.instance)
	if err != nil {
		return err
	}

	if len(devices) == 0 {
		return errors.New("failed to find GPUs with Vulkan support!")
	}

	fmt.Printf("Device count: %d\n", len(devices))

	var lastError error

	for _, device := range devices {
		suitable, err := this.isDeviceSuitable(device)
		if suitable {
			this.physicalDevice = device
			if err := this.deviceActualExtensions.fromDevice(device, ""); err != nil {
				lastError = err
			}
			break
		} else {
			lastError = err
		}
	}

	if this.physicalDevice == nil {
		if lastError != nil {
			return lastError
		} else {
			return errors.New("Failed to find a suitable GPU!")
		}
	}

	vk.GetPhysicalDeviceProperties(this.physicalDevice, &this.properties)

	fmt.Printf("Physical device: %s\n", this.properties.DeviceName)

	return nil
}
func (this *AxeDevice) createLogicalDevice() error {
	indices, err := this.findQueueFamilies(this.physicalDevice)
	if err != nil {
		return err
	}

	queueCreateInfos := []vk.DeviceQueueCreateInfo{}
	uniqueQueueFamilies := []uint32{indices.graphicsFamily}
	if indices.graphicsFamily != indices.presentFamily {
		uniqueQueueFamilies = append(uniqueQueueFamilies, indices.presentFamily)
	}

	for _, queueFamily := range uniqueQueueFamilies {
		queueCreateInfo := vk.DeviceQueueCreateInfo{}
		queueCreateInfo.QueueFamilyIndex = queueFamily
		queueCreateInfo.QueuePriorities = []float32{1.0}
		queueCreateInfos = append(queueCreateInfos, queueCreateInfo)
	}

	deviceFeatures := vk.PhysicalDeviceFeatures{}
	deviceFeatures.SamplerAnisotropy = true

	createInfo := vk.DeviceCreateInfo{}
	createInfo.QueueCreateInfos = queueCreateInfos
	createInfo.EnabledFeatures = &deviceFeatures
	createInfo.EnabledExtensionNames = this.deviceEnabledExtensions.getExtensions()

	// might not really be necessary anymore because device specific validation layers
	// have been deprecated
	if this.enableValidationLayers {
		createInfo.EnabledLayerNames = this.validationLayers.getExtensions()
	}

	if vk.CreateDevice(this.physicalDevice, &createInfo, nil, &this.device) != vk.SUCCESS {
		return errors.New("failed to create logical device!")
	}

	vk.GetDeviceQueue(this.device, indices.graphicsFamily, 0, &this.graphicsQueue)
	vk.GetDeviceQueue(this.device, indices.presentFamily, 0, &this.presentQueue)

	return nil
}
func (this *AxeDevice) createCommandPool() error {
	queueFamilyIndices, err := this.findPhysicalQueueFamilies()
	if err != nil {
		return err
	}

	poolInfo := vk.CommandPoolCreateInfo{}
	poolInfo.QueueFamilyIndex = queueFamilyIndices.graphicsFamily
	poolInfo.Flags = vk.CommandPoolCreateFlags(vk.COMMAND_POOL_CREATE_TRANSIENT_BIT | vk.COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT)

	if result := vk.CreateCommandPool(this.device, &poolInfo, nil, &this.commandPool); result != vk.SUCCESS {
		return VkError{result, "vkCreateCommandPool"}
	}
	return nil
}

func (this *AxeDevice) isDeviceSuitable(device vk.PhysicalDevice) (bool, error) {
	indices, err := this.findQueueFamilies(device)
	if err != nil {
		return false, err
	}

	fmt.Printf("Indices: %+v\n", indices)

	extensionsSupported := this.checkRequiredExtensionSupport(device)
	swapChainAdequate := false

	if extensionsSupported {
		swapChainSupport, err := this.querySwapChainSupport(device)
		if err != nil {
			return false, err
		}
		fmt.Printf("Swap chain support: %+v\n", swapChainSupport)
		swapChainAdequate = len(swapChainSupport.formats) > 0 && len(swapChainSupport.presentModes) > 0
	}

	fmt.Printf("Extensions supported: %+v\n", extensionsSupported)
	fmt.Printf("Swap chain adequate: %+v\n", swapChainAdequate)

	supportedFeatures := vk.PhysicalDeviceFeatures{}
	vk.GetPhysicalDeviceFeatures(device, &supportedFeatures)

	fmt.Printf("Supported features: %+v\n", supportedFeatures)

	return indices.isComplete() && extensionsSupported && swapChainAdequate && supportedFeatures.SamplerAnisotropy, nil
}
func (this *AxeDevice) getRequiredExtensions() instanceExtensions {
	extensions := this.instanceEnabledExtensions
	extensions.fromNames(this.window.window.GetRequiredInstanceExtensions())
	if this.enableValidationLayers {
		extensions.DebugUtils = true
	}
	return extensions
}
func (this *AxeDevice) checkValidationLayerSupport() (bool, error) {
	actualExtensions := layerExtensions{}
	err := actualExtensions.fromSystem()
	if err != nil {
		return false, err
	}

	return actualExtensions.isSupported(this.validationLayers), nil
}

func (this *AxeDevice) findQueueFamilies(device vk.PhysicalDevice) (QueueFamilyIndices, error) {
	indices := QueueFamilyIndices{}
	queueFamilies := getPhysicalDeviceQueueFamilyProperties(device)

	var i uint32 = 0
	for _, queueFamily := range queueFamilies {
		if queueFamily.QueueCount > 0 && (uint32(queueFamily.QueueFlags)&uint32(vk.QUEUE_GRAPHICS_BIT)) != 0 {
			indices.graphicsFamily = i
			indices.graphicsFamilyHasValue = true
		}
		presentSupport := false

		result := vk.ToGetPhysicalDeviceSurfaceSupportKHR(
			vk.GetInstanceProcAddr(this.instance, "vkGetPhysicalDeviceSurfaceSupportKHR"),
		)(device, i, this.surface, &presentSupport)
		if result != vk.SUCCESS {
			return indices, VkError{result, "vkGetPhysicalDeviceSurfaceSupportKHR"}
		}

		if queueFamily.QueueCount > 0 && presentSupport {
			indices.presentFamily = i
			indices.presentFamilyHasValue = true
		}
		if indices.isComplete() {
			break
		}
		i++
	}

	return indices, nil
}

var debugCallback vk.FuncDebugUtilsMessengerCallbackEXT = func(messageSeverity vk.DebugUtilsMessageSeverityFlagBitsEXT, messageType vk.DebugUtilsMessageTypeFlagsEXT, callbackData *vk.DebugUtilsMessengerCallbackDataEXT, userData unsafe.Pointer) bool {
	fmt.Printf("Debug Message %s\n", callbackData.Message)
	return false
}

//export printDebugMessage
func printDebugMessage(msg string) {
	fmt.Printf("Debug Message %s\n", msg)
}

func (this *AxeDevice) populateDebugMessengerCreateInfo(createInfo *vk.DebugUtilsMessengerCreateInfoEXT) {
	createInfo.MessageSeverity = vk.DebugUtilsMessageSeverityFlagsEXT(int(vk.DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT) | int(vk.DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT))
	createInfo.MessageType = vk.DebugUtilsMessageTypeFlagsEXT(int(vk.DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT) | int(vk.DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT) | int(vk.DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT))
	createInfo.UserData = nil // Optional
	createInfo.UserCallback = vk.PFNDebugUtilsMessengerCallbackEXT(unsafe.Pointer(&debugCallback))
}

func (this *AxeDevice) hasGflwRequiredInstanceExtensions() error {
	requiredExtensions := this.getRequiredExtensions()
	missing := this.instanceAvailableExtensions.getMissing(requiredExtensions)
	if len(missing) > 0 {
		return fmt.Errorf("Missing required glfw extensions: %v", missing)
	}

	return nil
}
func (this *AxeDevice) checkRequiredExtensionSupport(device vk.PhysicalDevice) bool {
	available := deviceExtensions{}
	available.fromDevice(device, "")

	missing := available.getMissing(this.deviceRequiredExtensions)

	if len(missing) > 0 {
		fmt.Printf("Missing extensions: %v\n", missing)
	}

	return len(missing) == 0
}
func (this *AxeDevice) querySwapChainSupport(device vk.PhysicalDevice) (SwapChainSupportDetails, error) {
	details := SwapChainSupportDetails{}

	if this.instanceAvailableExtensions.PhysicalDeviceProperties && this.instanceAvailableExtensions.Surface {
		result := vk.ToGetPhysicalDeviceSurfaceCapabilitiesKHR(
			vk.GetInstanceProcAddr(this.instance, "vkGetPhysicalDeviceSurfaceCapabilitiesKHR"),
		)(device, this.surface, &details.capabilities)

		if result != vk.SUCCESS {
			return details, VkError{result, "vkGetPhysicalDeviceSurfaceCapabilitiesKHR"}
		}

		var err error

		details.formats, err = GetSliceError(func(count *uint32, out []vk.SurfaceFormatKHR) error {
			return checkResult(
				vk.ToGetPhysicalDeviceSurfaceFormatsKHR(
					vk.GetInstanceProcAddr(this.instance, "vkGetPhysicalDeviceSurfaceFormatsKHR"),
				)(device, this.surface, count, out),
				"vkGetPhysicalDeviceSurfaceFormatsKHR",
			)
		})
		if err != nil {
			return details, err
		}

		details.presentModes, err = GetSliceError(func(count *uint32, out []vk.PresentModeKHR) error {
			return checkResult(
				vk.ToGetPhysicalDeviceSurfacePresentModesKHR(
					vk.GetInstanceProcAddr(this.instance, "vkGetPhysicalDeviceSurfacePresentModesKHR"),
				)(device, this.surface, count, out),
				"vkGetPhysicalDeviceSurfacePresentModesKHR",
			)
		})
		if err != nil {
			return details, err
		}
	}

	return details, nil
}

func (this *AxeDevice) getSwapChainSupport() (SwapChainSupportDetails, error) {
	return this.querySwapChainSupport(this.physicalDevice)
}
func (this *AxeDevice) findMemoryType(typeFilter uint32, properties vk.MemoryPropertyFlags) (uint32, error) {
	memProperties := vk.PhysicalDeviceMemoryProperties{}
	vk.GetPhysicalDeviceMemoryProperties(this.physicalDevice, &memProperties)
	var i uint32
	for i = 0; i < memProperties.MemoryTypeCount; i++ {
		if (typeFilter&(1<<i)) != 0 && (uint32(memProperties.MemoryTypes[i].PropertyFlags)&uint32(properties)) == uint32(properties) {
			return i, nil
		}
	}
	return 0, errors.New("There was a problem finding a suitable memory type.")
}
func (this *AxeDevice) findPhysicalQueueFamilies() (QueueFamilyIndices, error) {
	return this.findQueueFamilies(this.physicalDevice)
}
func (this *AxeDevice) findSupportedFormat(candidates []vk.Format, tiling vk.ImageTiling, features vk.FormatFeatureFlags) (vk.Format, error) {
	for _, format := range candidates {
		props := vk.FormatProperties{}
		vk.GetPhysicalDeviceFormatProperties(this.physicalDevice, format, &props)

		if tiling == vk.IMAGE_TILING_LINEAR && (props.LinearTilingFeatures&features) == features {
			return format, nil
		} else if tiling == vk.IMAGE_TILING_OPTIMAL && (props.OptimalTilingFeatures&features) == features {
			return format, nil
		}
	}
	return 0, errors.New("failed to find supported format!")
}

// Buffer Helper Functions
func (this *AxeDevice) createBuffer(size vk.DeviceSize, usage vk.BufferUsageFlags, properties vk.MemoryPropertyFlags, buffer *vk.Buffer, bufferMemory *vk.DeviceMemory) error {
	bufferInfo := vk.BufferCreateInfo{}
	bufferInfo.Size = size
	bufferInfo.Usage = usage
	bufferInfo.SharingMode = vk.SHARING_MODE_EXCLUSIVE

	if result := vk.CreateBuffer(this.device, &bufferInfo, nil, buffer); result != vk.SUCCESS {
		return VkError{result, "vkCreateBuffer"}
	}

	memRequirements := vk.MemoryRequirements{}
	vk.GetBufferMemoryRequirements(this.device, *buffer, &memRequirements)

	allocInfo := vk.MemoryAllocateInfo{}
	allocInfo.AllocationSize = memRequirements.Size
	allocInfo.MemoryTypeIndex, _ = this.findMemoryType(memRequirements.MemoryTypeBits, properties)

	if result := vk.AllocateMemory(this.device, &allocInfo, nil, bufferMemory); result != vk.SUCCESS {
		return VkError{result, "vkAllocateMemory"}
	}

	if result := vk.BindBufferMemory(this.device, *buffer, *bufferMemory, 0); result != vk.SUCCESS {
		return VkError{result, "vkBindBufferMemory"}
	}
	return nil
}
func (this *AxeDevice) beginSingleTimeCommands() (vk.CommandBuffer, error) {
	allocInfo := vk.CommandBufferAllocateInfo{}
	allocInfo.Level = vk.COMMAND_BUFFER_LEVEL_PRIMARY
	allocInfo.CommandPool = this.commandPool
	allocInfo.CommandBufferCount = 1

	commandBuffer := make([]vk.CommandBuffer, 1)
	if result := vk.AllocateCommandBuffers(this.device, &allocInfo, commandBuffer); result != vk.SUCCESS {
		return nil, VkError{result, "vkAllocateCommandBuffers"}
	}

	beginInfo := vk.CommandBufferBeginInfo{}
	beginInfo.Flags = vk.CommandBufferUsageFlags(vk.COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT)

	if result := vk.BeginCommandBuffer(commandBuffer[0], &beginInfo); result != vk.SUCCESS {
		return nil, VkError{result, "vkBeginCommandBuffer"}
	}

	return commandBuffer[0], nil
}
func (this *AxeDevice) endSingleTimeCommands(commandBuffer vk.CommandBuffer) error {
	if result := vk.EndCommandBuffer(commandBuffer); result != vk.SUCCESS {
		return VkError{result, "vkEndCommandBuffer"}
	}

	commandBuffers := []vk.CommandBuffer{commandBuffer}
	submitInfo := make([]vk.SubmitInfo, 1)
	submitInfo[0].CommandBuffers = commandBuffers

	if result := vk.QueueSubmit(this.graphicsQueue, submitInfo, 0); result != vk.SUCCESS {
		return VkError{result, "vkQueueSubmit"}
	}
	if result := vk.QueueWaitIdle(this.graphicsQueue); result != vk.SUCCESS {
		return VkError{result, "vkQueueWaitIdle"}
	}

	vk.FreeCommandBuffers(this.device, this.commandPool, commandBuffers)

	return nil
}
func (this *AxeDevice) copyBuffer(srcBuffer vk.Buffer, dstBuffer vk.Buffer, size vk.DeviceSize) error {
	commandBuffer, err := this.beginSingleTimeCommands()
	if err != nil {
		return err
	}

	copyRegion := make([]vk.BufferCopy, 1)
	copyRegion[0].SrcOffset = 0 // Optional
	copyRegion[0].DstOffset = 0 // Optional
	copyRegion[0].Size = size
	vk.CmdCopyBuffer(commandBuffer, srcBuffer, dstBuffer, copyRegion)

	err = this.endSingleTimeCommands(commandBuffer)
	if err != nil {
		return err
	}
	return nil
}
func (this *AxeDevice) copyBufferToImage(buffer vk.Buffer, image vk.Image, width uint32, height uint32, layerCount uint32) error {
	commandBuffer, err := this.beginSingleTimeCommands()
	if err != nil {
		return err
	}

	region := make([]vk.BufferImageCopy, 1)
	region[0].BufferOffset = 0
	region[0].BufferRowLength = 0
	region[0].BufferImageHeight = 0
	region[0].ImageSubresource.AspectMask = vk.ImageAspectFlags(vk.IMAGE_ASPECT_COLOR_BIT)
	region[0].ImageSubresource.MipLevel = 0
	region[0].ImageSubresource.BaseArrayLayer = 0
	region[0].ImageSubresource.LayerCount = layerCount
	region[0].ImageOffset = vk.Offset3D{X: 0, Y: 0, Z: 0}
	region[0].ImageExtent = vk.Extent3D{Width: width, Height: height, Depth: 1}

	vk.CmdCopyBufferToImage(commandBuffer, buffer, image, vk.IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL, region)

	err = this.endSingleTimeCommands(commandBuffer)
	if err != nil {
		return err
	}
	return nil
}
func (this *AxeDevice) createImageWithInfo(imageInfo *vk.ImageCreateInfo, properties vk.MemoryPropertyFlags, image *vk.Image, imageMemory vk.DeviceMemory) error {
	if vk.CreateImage(this.device, imageInfo, nil, image) != vk.SUCCESS {
		return errors.New("Failed to create image.")
	}

	memRequirements := vk.MemoryRequirements{}
	vk.GetImageMemoryRequirements(this.device, *image, &memRequirements)

	allocInfo := vk.MemoryAllocateInfo{}
	allocInfo.AllocationSize = memRequirements.Size
	allocInfo.MemoryTypeIndex, _ = this.findMemoryType(memRequirements.MemoryTypeBits, properties)

	if vk.AllocateMemory(this.device, &allocInfo, nil, &imageMemory) != vk.SUCCESS {
		return errors.New("Failed to allocate image memory.")
	}

	if vk.BindImageMemory(this.device, *image, imageMemory, 0) != vk.SUCCESS {
		return errors.New("Failed to bind image memory.")
	}

	return nil
}

type AxeWindow struct {
	window              *glfw.Window
	windowName          string
	width               int
	height              int
	framebufferResized  bool
	framebufferCallback glfw.FramebufferSizeCallback
}

func newWindow(width int, height int, name string) (*AxeWindow, error) {
	win := new(AxeWindow)
	win.width = width
	win.height = height
	win.windowName = name
	win.framebufferCallback = func(w *glfw.Window, width, height int) {
		win.framebufferResized = true
		win.width = width
		win.height = height

		fmt.Printf("frame resize %v %v\n", width, height)
	}
	err := win.initWindow()
	return win, err
}

func (this *AxeWindow) destroy() {
	this.window.Destroy()
	glfw.Terminate()
}
func (this *AxeWindow) initWindow() error {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	if !glfw.VulkanSupported() {
		panic("Vulkan is not supported with your version of GLFW.")
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.Resizable, glfw.True)

	this.window, err = glfw.CreateWindow(this.width, this.height, this.windowName, nil, nil)
	if err != nil {
		return err
	}
	// this.window.MakeContextCurrent()
	this.window.SetFramebufferSizeCallback(this.framebufferCallback)

	return nil
}
func (this *AxeWindow) shouldClose() bool {
	return this.window.ShouldClose()
}
func (this *AxeWindow) getExtent() vk.Extent2D {
	return vk.Extent2D{Width: uint32(this.width), Height: uint32(this.height)}
}
func (this *AxeWindow) wasWindowResized() bool {
	return this.framebufferResized
}
func (this *AxeWindow) resetWindowResizedFlag() {
	this.framebufferResized = false
}
func (this *AxeWindow) getGLFWwindow() *glfw.Window {
	return this.window
}
func (this *AxeWindow) createWindowSurface(instance vk.Instance) (vk.SurfaceKHR, error) {
	surface, err := this.window.CreateWindowSurface(instance, nil)
	// surfaceRef := *(*vk.SurfaceKHR)(unsafe.Pointer(surface))
	// fmt.Printf("Surface: %v\n", surfaceRef)
	return vk.SurfaceKHR(surface), err
}

func getInstanceLayerProperties() ([]vk.LayerProperties, error) {
	return GetSliceError(func(count *uint32, out []vk.LayerProperties) error {
		return checkResult(vk.EnumerateInstanceLayerProperties(count, out), "There was a problem")
	})
}

func getPhysicalDeviceQueueFamilyProperties(physicalDevice vk.PhysicalDevice) []vk.QueueFamilyProperties {
	return GetSlice(func(count *uint32, out []vk.QueueFamilyProperties) {
		vk.GetPhysicalDeviceQueueFamilyProperties(physicalDevice, count, out)
	})
}

func getPhysicalDevices(instance vk.Instance) ([]vk.PhysicalDevice, error) {
	return GetSliceError(func(count *uint32, out []vk.PhysicalDevice) error {
		return checkResult(vk.EnumeratePhysicalDevices(instance, count, out), "vkEnumeratePhysicalDevices")
	})
}

func getInstanceExtensionProperties(layerName string) ([]vk.ExtensionProperties, error) {
	return GetSliceError(func(count *uint32, out []vk.ExtensionProperties) error {
		return checkResult(vk.EnumerateInstanceExtensionProperties(layerName, count, out), "vkEnumerateInstanceExtensionProperties")
	})
}

func getDeviceExtensionProperties(physicalDevice vk.PhysicalDevice, layerName string) ([]vk.ExtensionProperties, error) {
	return GetSliceError(func(count *uint32, out []vk.ExtensionProperties) error {
		return checkResult(vk.EnumerateDeviceExtensionProperties(physicalDevice, layerName, count, out), "vkEnumerateDeviceExtensionProperties")
	})
}

func setFromNames(out any, names []string) {
	fieldMap := make(map[string]reflect.Value)
	outReflect := reflect.ValueOf(out).Elem()
	outType := outReflect.Type()
	if outType.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < outReflect.NumField(); i++ {
		outField := outType.Field(i)
		if name, ok := outField.Tag.Lookup("name"); ok {
			key := strings.ToLower(name)
			fieldMap[key] = outReflect.Field(i)
		}
	}

	for _, name := range names {
		key := strings.ToLower(name)
		if value, ok := fieldMap[key]; ok && value.Kind() == reflect.Bool {
			value.SetBool(true)
		}
	}
}

func getMissingNames(actualValue any, expectedValue any) []string {
	missing := make([]string, 0)
	actual := reflect.ValueOf(actualValue)
	expect := reflect.ValueOf(expectedValue)
	for i := 0; i < expect.NumField(); i++ {
		actualValue := actual.Field(i)
		expectValue := expect.Field(i)
		if expectValue.Bool() && !actualValue.Bool() {
			field := actual.Type().Field(i)
			if ext, ok := field.Tag.Lookup("name"); ok {
				missing = append(missing, ext)
			}
		}
	}
	return missing
}

func GetSlice[T any](getter func(*uint32, []T)) []T {
	var count uint32 = 0
	getter(&count, nil)
	out := make([]T, count)
	getter(&count, out)
	return out
}

func GetSliceError[T any](getter func(*uint32, []T) error) ([]T, error) {
	var count uint32 = 0
	err := getter(&count, nil)
	if err != nil {
		return []T{}, err
	}
	out := make([]T, count)
	err = getter(&count, out)
	return out, err
}

func MapSlice[S any, D any](source []S, mapper func(source S) D) []D {
	dest := make([]D, len(source))
	for i, item := range source {
		dest[i] = mapper(item)
	}
	return dest
}

func checkResult(result vk.Result, functionName string) error {
	if result != vk.SUCCESS {
		return VkError{result, functionName}
	}
	return nil
}

type VkError struct {
	Result vk.Result
	Func   string
}

var _ error = &VkError{}

func (e VkError) Error() string {
	return fmt.Sprintf("There was an error calling '%s', result: %v.", e.Func, e.Result)
}

package api

// /// Something that provides a runtime api.
// pub trait ProvideRuntimeApi<Block: BlockT> {
type ProvideRuntimeAPI[API any] interface {
	// 	/// The concrete type that provides the api.
	// 	type Api: ApiExt<Block>;

	// /// Returns the runtime api.
	// /// The returned instance will keep track of modifications to the storage. Any successful
	// /// call to an api function, will `commit` its changes to an internal buffer. Otherwise,
	// /// the modifications will be `discarded`. The modifications will not be applied to the
	// /// storage, even on a `commit`.
	// fn runtime_api(&self) -> ApiRef<Self::Api>;
	RuntimeAPI() API
}

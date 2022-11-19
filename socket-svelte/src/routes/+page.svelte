<script>
	import { onMount } from 'svelte';
  /**
	 * @type {WebSocket | null}
	 */
  let socket = null
	const handleClick = () => {
    socket?.send(JSON.stringify({"greeting": value}))
	};


	const onMessage = (/** @type {string} */ e) => {
		messagesList = messagesList.concat(e)
	}
	onMount(() => {
		socket = new WebSocket('ws://localhost:8000/socket');
		socket.onmessage = (e) => {
			console.log('got message', e.data)
			onMessage(e.data)
		}
		messagesList = 'hello'
		});

	$: value = '';
	$: messagesList = ''
</script>

<div class="flex flex-col justify-center items-center bg-slate-500 h-screen gap-4">
  <div class="bg-white rounded-md p-4">
    <input class="p-4 bg-slate-300" type="text" bind:value placeholder="write something" />
    <button class="bg-blue-500 text-white p-4 rounded-md" on:click={handleClick}>send</button>
  </div>
	<div class="p-4 bg-white rounded-md">
		<textarea
			bind:value={messagesList}
			class="" name="messages" id="" cols="30" rows="10" placeholder="messages go here"></textarea>
	</div>
</div>
